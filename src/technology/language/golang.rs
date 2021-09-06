//! Golang specific tools.

use serde::{Deserialize, Serialize};
use std::{collections::HashMap, path::Path};
use tera::{Result, Tera, Value};

use crate::{
    model::{data_type::DataType, Model},
    render::render_simple,
    technology::Technology,
};

/// "imports" for different golang files
#[derive(Serialize, Deserialize, Debug)]
pub struct Imports {
    /// some data_types may require specific imports in "model.go"
    pub model: Vec<&'static str>,
    /// each database has its own library
    pub database_library: &'static str,
}

/// Root of the golang struct
#[derive(Serialize, Deserialize, Debug)]
pub struct Golang {
    pub imports: Imports,
    /// db_driver is used as the param of `sql.Open`
    pub db_driver: &'static str,
}

// TODO: It is highly possible that each language needs `data_type`, `register` and `render`, maybe we can add a trait.
/// Get the string representing of a data_type in Golang.
fn data_type(data_type: DataType) -> String {
    match data_type {
        DataType::Int(x) if x <= 8 => "int8".to_string(),
        DataType::Int(x) if x <= 16 => "int16".to_string(),
        DataType::Int(x) if x <= 32 => "int32".to_string(),
        DataType::Int(_) => "int64".to_string(),
        DataType::UInt(x) if x <= 8 => "uint8".to_string(),
        DataType::UInt(x) if x <= 16 => "uint16".to_string(),
        DataType::UInt(x) if x <= 32 => "uint32".to_string(),
        DataType::UInt(_) => "uint64".to_string(),
        DataType::Float(x) if x <= 32 => "float32".to_string(),
        DataType::Float(_) => "float64".to_string(),
        DataType::String(_) => "string".to_string(),
        DataType::DateTime => "time".to_string(),
    }
}

fn data_type_in_template(args: &HashMap<String, Value>) -> Result<Value> {
    let data_type_value: DataType = serde_json::from_value(args.get("data_type").unwrap().clone())?;
    Ok(Value::String(data_type(data_type_value)))
}

/// A function which can be used in the template for judging whether the datatype is a string
/// Useful when accepting user input. See `templates/go/handler.go.template` for detail.
fn data_type_is_string(args: &HashMap<String, Value>) -> Result<Value> {
    let data_type_value: DataType =
        serde_json::from_value(args.get("data_type").unwrap_or(&Value::Null).clone())?;
    Ok(Value::Bool(matches!(data_type_value, DataType::String(_))))
}

impl Golang {
    pub fn new(technology: &Technology, model: &Model) -> Self {
        let imports_model = if model
            .fields
            .iter()
            .any(|field| field.data_type == DataType::DateTime)
        {
            vec!["time"]
        } else {
            vec![]
        };
        let (db_driver, imports_database) = match technology.database {
            crate::technology::DataBase::PgSQL => ("pgsql", "github.com/lib/pq"),
            crate::technology::DataBase::MySQL => ("mysql", "github.com/go-sql-driver/mysql"),
        };
        Self {
            imports: Imports {
                model: imports_model,
                database_library: imports_database,
            },
            db_driver,
        }
    }
}

pub fn register(
    tera: &mut Tera,
    technology: &Technology,
    model: &Model,
    context: &mut tera::Context,
) {
    tera.register_function("data_type", data_type_in_template);
    tera.register_function("is_string", data_type_is_string);
    let golang = Golang::new(technology, model);
    context.insert("golang", &golang);
}

pub fn render(tera: &Tera, home: impl AsRef<Path>, context: &mut tera::Context) {
    render_simple(tera, home.as_ref(), "go", "go.mod", context);
    render_simple(tera, home.as_ref(), "go", "model/model.go", context);
    render_simple(tera, home.as_ref(), "go", "infrastructure/db.go", context);
    render_simple(tera, home.as_ref(), "go", "handler/handler.go", context);
    render_simple(tera, home, "go", "main.go", context);
}
