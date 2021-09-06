mod database;
pub mod language;
pub use database::DataBase;

// TODO: language enum
/// Technology is about the specific technology the user want to implement the API.
pub struct Technology {
    pub database: DataBase,
}
