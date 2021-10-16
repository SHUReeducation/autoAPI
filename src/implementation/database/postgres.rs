//! PostgreSQL database tool functions
use std::collections::HashMap;

use tera::{Result, Value};

use crate::model::data_type::DataType;

fn db_type(data_type: DataType) -> String {
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
        DataType::String(None) => "text".to_string(),
        DataType::String(Some(x)) if x.fixed_length => format!("char({})", x.length),
        DataType::String(Some(x)) if !x.fixed_length => format!("varchar({})", x.length),
        DataType::String(Some(_)) => unreachable!(),
        DataType::DateTime => "time".to_string(),
    }
}

fn db_type_in_template(args: &HashMap<String, Value>) -> Result<Value> {
    let data_type_value: DataType = serde_json::from_value(args.get("data_type").unwrap().clone())?;
    Ok(Value::String(db_type(data_type_value)))
}

pub fn register(tera: &mut tera::Tera) {
    tera.register_function("db_type", db_type_in_template);
}
