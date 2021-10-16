use std::fs::OpenOptions;

use anyhow::{anyhow, Result};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Field {
    pub name: String,
    #[serde(rename = "type")]
    pub data_type: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct API {
    pub name: String,
    pub fields: Vec<Field>,
}

#[derive(Serialize, Deserialize, Debug, Clone, Default)]
pub struct DataBase {
    pub engine: Option<String>,
    #[serde(alias = "address")]
    #[serde(alias = "connection_string")]
    pub url: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Implementation {
    pub framework: String,
    #[serde(default)]
    #[serde(alias = "db")]
    pub database: DataBase,
}

fn fn_true() -> bool {
    true
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Docker {
    #[serde(default = "fn_true")]
    pub generate_file: bool,
    pub username: Option<String>,
    pub tag: Option<String>,
}

impl Default for Docker {
    fn default() -> Self {
        Self {
            generate_file: true,
            username: None,
            tag: None,
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct CICD {
    #[serde(default)]
    pub docker: Docker,
    #[serde(alias = "k8s")]
    pub kubernetes: Option<bool>,
    #[serde(default = "fn_true")]
    #[serde(alias = "ghaction")]
    #[serde(alias = "gh_action")]
    pub github_action: bool,
}

impl Default for CICD {
    fn default() -> Self {
        Self {
            docker: Default::default(),
            kubernetes: None,
            github_action: true,
        }
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Config {
    pub api: Option<API>,
    pub implementation: Implementation,
    #[serde(default)]
    pub cicd: CICD,
}

