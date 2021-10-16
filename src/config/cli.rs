use std::path::PathBuf;
use structopt::StructOpt;

#[derive(Debug, StructOpt)]
#[structopt(
    name = "autoapi",
    about = "A tool for generating CRUD API program automatically.",
    rename_all = "kebab"
)]
pub struct Config {
    /// Config file path.
    #[structopt(short="i", long, alias="input", parse(from_os_str), env = "CONFIG")]
    pub config: PathBuf,
    /// Output project path.
    #[structopt(short, long, parse(from_os_str), env = "OUTPUT")]
    pub output: PathBuf,
    /// Set this flag to overwrite the `output` directory before generating instead of report an error.
    #[structopt(short, long, env = "FORCE")]
    pub force: bool,

    /// Output project path.
    #[structopt(alias="dbms", long, env = "DATABASE")]
    pub database_engine: Option<String>,
    #[structopt(short, long, env = "API_NAME")]
    pub name: Option<String>,

    #[structopt(alias="ddl", long, env = "DDL")]
    pub load_from_ddl: Option<String>,
    #[structopt(short="load-db", long, env = "LOAD_DB")]
    pub load_from_db: Option<String>,

    #[structopt(short="d", long, env = "DOCKER")]
    pub generate_docker: Option<bool>, 
    #[structopt(alias="du", long, env = "DOCKER_USERNAME")]
    pub docker_username: Option<String>,
    #[structopt(alias="dt", long, env = "DOCKER_TAG")]
    pub docker_tag: Option<String>, 

    #[structopt(short="k8s", long, env = "KUBERNETES")]
    pub kubernetes: Option<bool>,
}
