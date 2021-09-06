//! This module contains the "model" related code.
//! "model" here means how the user wants the API to look like,
//! but has nothing to do with the any specific technology the user selects.

use serde::{Deserialize, Serialize};

use self::data_type::DataType;

pub mod data_type;

/// Field means a field of a model.
/// Can usually be mapped to a database column.
#[derive(Debug, Serialize, Deserialize)]
pub struct Field {
    pub name: String,
    pub data_type: DataType,
}

/// Model is the object the user wants to generate API for.
/// Can usually be mapped to a database table.
#[derive(Debug, Serialize, Deserialize)]
pub struct Model {
    pub name: String,
    pub primary_key: Field,
    pub fields: Vec<Field>,
}
