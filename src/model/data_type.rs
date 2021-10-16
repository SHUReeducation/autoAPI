use convert_case::{Case, Casing};
use serde::{Deserialize, Serialize};

/// Meta information about a string.
#[derive(Serialize, Deserialize, Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
pub struct StringMeta {
    pub length: usize,
    /// If a string's length is fixed, we can use `CHAR` type in database.
    pub fixed_length: bool,
}

/// A type of data.
#[derive(Serialize, Deserialize, Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
#[serde(rename_all = "lowercase")]
pub enum DataType {
    Int(usize),
    UInt(usize),
    Float(usize),
    String(Option<StringMeta>),
    // TODO: separate date, time and datetime.
    DateTime,
}

impl From<String> for DataType {
    fn from(s: String) -> Self {
        match s.as_str().to_case(Case::Snake).as_str() {
            "tiny_int" | "tinyint" | "tiny_integer" | "int8" => DataType::Int(8),
            "small_int" | "smallint" | "small_integer" | "int16" => DataType::Int(16),
            "int" | "integer" | "medium_int" | "mediumint" | "medium_integer" | "int32" => {
                DataType::Int(32)
            }
            "big_int" | "bigint" | "big_integer" | "int64" => DataType::Int(64),
            "unsigned_tiny_int" | "unsigned_tiny_integer" | "uint8" => DataType::UInt(8),
            "unsigned_small_int" | "unsigned_small_integer" | "uint16" => DataType::UInt(16),
            "uint" | "unsigned" | "unsigned_medium_int" | "unsigned_medium_integer" | "uint32" => {
                DataType::UInt(32)
            }
            "unsigned_big_int" | "unsigned_big_integer" | "uint64" => DataType::UInt(64),
            "float" | "float32" => DataType::Float(32),
            "double" | "float64" => DataType::Float(64),
            "string" | "text" | "medium_text" => DataType::String(None),
            s => {
                if s.starts_with("char(") {
                    let length: usize = s
                        .trim_start_matches("char(")
                        .trim_end_matches(')')
                        .parse()
                        .unwrap();
                    DataType::String(Some(StringMeta {
                        length,
                        fixed_length: true,
                    }))
                } else if s.starts_with("varchar(") {
                    let length: usize = s
                        .trim_start_matches("varchar(")
                        .trim_end_matches(')')
                        .parse()
                        .unwrap();
                    DataType::String(Some(StringMeta {
                        length,
                        fixed_length: false,
                    }))
                } else {
                    todo!(
                        "will support these in the future:
                    tiny_text,
                    long_text,
                    tiny_blob,
                    blob,
                    medium_blob,
                    long_blob,
                    bool,
                    date,
                    date_time,
                    time,
                    timestamp"
                    )
                }
            }
        }
    }
}
