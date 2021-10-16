use implementation::Implementation;

use crate::{
    implementation::framework::golang,
    model::{data_type::DataType, Field, Model},
};

mod config;
mod implementation;
mod model;
mod render;

fn main() {
    let mut tera = render::load_templates();
    render::filter::register(&mut tera);
    let config = config::from_cli_config().unwrap();
    let implementation = Implementation {
        database: implementation::DataBase::PgSQL,
    };
    let model = Model {
        name: config.generate_config.api.as_ref().unwrap().name.clone(),
        primary_key: Field {
            name: "id".to_string(),
            data_type: DataType::UInt(64),
        },
        fields: config
            .generate_config
            .api
            .unwrap()
            .fields
            .into_iter()
            .map(|it| Field {
                name: it.name,
                data_type: it.data_type.into(),
            })
            .collect(),
    };
    let mut context = tera::Context::new();
    context.insert("model", &model);
    golang::register(&mut tera, &implementation, &model, &mut context);
    golang::render(&tera, config.output, &mut context);
}
