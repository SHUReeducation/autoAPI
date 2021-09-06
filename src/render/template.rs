use include_dir::{include_dir, Dir, DirEntry};
use tera::Tera;

/// All templates are included into the binary, so the binary can be distributed alone without any folders.
static TEMPLATES_DIR: Dir = include_dir!("./templates");

/// Loads all templates from the templates "directory".
pub fn load_templates() -> Tera {
    let mut tera = Tera::default();
    for entry in TEMPLATES_DIR.find("**/*.template").unwrap() {
        if let DirEntry::File(f) = entry {
            tera.add_raw_template(f.path().to_str().unwrap(), f.contents_utf8().unwrap())
                .unwrap();
        }
    }
    tera
}
