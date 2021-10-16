mod cli;
mod file;
use anyhow::anyhow;
use std::{fs::File, io::Read, path::PathBuf};

use structopt::StructOpt;

pub struct Config {
    pub output: PathBuf,
    pub force: bool,
    // For now, `file::Config` contains all the configuration options we need for generating the API.
    pub generate_config: file::Config,
}

pub fn from_cli_config() -> anyhow::Result<Config> {
    let cli_config = cli::Config::from_args_safe()?;
    let mut config_file = File::open(&cli_config.config)?;
    let mut file_config: file::Config = match cli_config
        .config
        .extension()
        .and_then(|it| it.to_str())
        .ok_or(anyhow!("Cannot open config file"))?
    {
        "toml" => {
            let mut content = String::new();
            config_file.read_to_string(&mut content)?;
            toml::from_str(&content)?
        }
        "json" => serde_json::from_reader(config_file)?,
        "yaml" => serde_yaml::from_reader(config_file)?,
        _ => return Err(anyhow!("Unsupported config file type")),
    };
    // Merging config from file and cli
    if file_config.implementation.database.engine.is_none() {
        file_config.implementation.database.engine = Some("pgsql".to_string());
    }
    if let Some(database_engine) = cli_config.database_engine {
        file_config.implementation.database.engine = Some(database_engine);
    }
    if file_config.implementation.database.engine.is_none()
        && file_config.implementation.database.url.is_some()
    {
        let mut url = file_config.implementation.database.url.take().unwrap();
        if url.starts_with("postgres://") {
            file_config.implementation.database.engine = Some("postgres".to_string());
        } else if url.starts_with("pgsql://") {
            println!("warning: pgsql:// won't work, I'll use postgres:// instead");
            url = url.replace("pgsql://", "postgres://");
            file_config.implementation.database.engine = Some("postgres".to_string());
        } else if url.starts_with("mysql://") {
            file_config.implementation.database.engine = Some("mysql".to_string());
        } else if url.starts_with("sqlite://") {
            file_config.implementation.database.engine = Some("sqlite".to_string());
        } else {
            return Err(anyhow!(
                "database engine not set and cannot auto refereed from connection string"
            ));
        }
        file_config.implementation.database.url = Some(url);
    }
    if file_config.api.is_none() {
        if let Some(_load_from_ddl) = cli_config.load_from_ddl {
            todo!("load api config from ddl_file");
        }
        let try_load_from_db = if let Some(addr) = cli_config.load_from_db {
            Some(addr)
        } else if let Some(addr) = file_config.implementation.database.url {
            Some(addr)
        } else {
            None
        };
        if try_load_from_db.is_some() {
            if cli_config.name.is_none() {
                return Err(anyhow!("I need to know API name before load it from db!"));
            }
        }
        todo!("load api config from database");
    }

    if let Some(generate_docker) = cli_config.generate_docker {
        file_config.cicd.docker.generate_file = generate_docker;
    }
    if let Some(docker_tag) = cli_config.docker_tag {
        file_config.cicd.docker.tag = Some(docker_tag);
    }
    if let Some(docker_username) = cli_config.docker_username {
        file_config.cicd.docker.username = Some(docker_username);
    }
    if let Some(k8s) = cli_config.kubernetes {
        file_config.cicd.kubernetes = Some(k8s);
    }
    file_config.cicd.kubernetes = match file_config.cicd.kubernetes {
        Some(false) => Some(false),
        Some(true) if file_config.implementation.database.url.is_none() => {
            return Err(anyhow!("generating kubernetes yaml file requires database connection string"));
        }
        Some(true) if file_config.cicd.docker.username.is_none() => {
            // todo: support other docker image registry than dockerhub
            return Err(anyhow!("generating kubernetes yaml file requires docker username"));
        }
        Some(true) => Some(true),
        None if file_config.implementation.database.url.is_none() => Some(false),
        None => Some(true)
    };
    Ok(Config {
        output: cli_config.output,
        force: cli_config.force,
        generate_config: file_config,
    })
}
