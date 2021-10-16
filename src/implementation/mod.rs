mod database;
pub mod framework;
pub use database::DataBase;

// TODO: framework enum
/// Technology is about the specific technology the user want to implement the API.
pub struct Implementation {
    pub database: DataBase,
}
