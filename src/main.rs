use technology::Technology;

use crate::{
    model::{data_type::DataType, Field, Model},
    technology::language::golang,
};

mod model;
mod render;
mod technology;

fn main() {
    let mut tera = render::load_templates();
    render::filter::register(&mut tera);
    let config = Technology {
        database: technology::DataBase::PgSQL,
    };
    let model = Model {
        name: "shuSB".to_string(),
        primary_key: Field {
            name: "id".to_string(),
            data_type: DataType::UInt(64),
        },
        fields: vec![
            Field {
                name: "name".to_string(),
                data_type: DataType::String(None),
            },
            Field {
                name: "IQ".to_string(),
                data_type: DataType::Int(32),
            },
        ],
    };
    let mut context = tera::Context::new();
    context.insert("model", &model);
    golang::register(&mut tera, &config, &model, &mut context);
    golang::render(&tera, "./shuSB", &mut context);
}
