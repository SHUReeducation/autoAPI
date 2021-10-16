//! Database the user wants to use and related tools.

use serde::{Deserialize, Serialize};

/// The database the user wants to use.
#[derive(Serialize, Deserialize, Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
#[serde(rename_all = "lowercase")]
pub enum DataBase {
    PgSQL,
    MySQL,
}
