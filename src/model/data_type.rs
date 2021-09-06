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
