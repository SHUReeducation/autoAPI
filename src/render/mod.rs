//! Template rendering utils shared between technologies.
//! Note: Currently, for each language, call <language>::render for the actual rendering.

pub mod filter;
mod template;
use std::{
    fs::{self, File},
    io::Write,
    path::Path,
};
pub use template::load_templates;

use tera::Tera;

/// Render home/filename with context.
pub fn render_simple(
    tera: &Tera,
    home: impl AsRef<Path>,
    language: &str,
    filename: &str,
    context: &tera::Context,
) {
    let path_str = format!("{}/{}", home.as_ref().display(), filename);
    let path = Path::new(&path_str);
    let parent_path = path.parent().unwrap();
    fs::create_dir_all(parent_path).unwrap();
    let mut f = File::create(path).unwrap();
    let template_name = format!("{}/{}.template", language, filename);
    let content = tera.render(&template_name, context).unwrap();
    f.write_all(content.as_bytes()).unwrap();
}
