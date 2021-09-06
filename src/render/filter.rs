//! Useful filters to be registered into [tera](https://tera.netlify.app/docs/)'s template.

use std::collections::HashMap;

use convert_case::{Case, Casing};
use pluralize_rs::{to_plural, to_singular};
use tera::{Error, Result, Tera, Value};

/// Change the case of a string into camelCase.
fn camel_case(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    Ok(Value::String(
        value
            .as_str()
            .ok_or_else(|| Error::msg("camel_case filter can only applied on strings"))?
            .to_case(Case::Camel),
    ))
}

/// Change the case of a string into kebab-case.
fn kebab_case(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    Ok(Value::String(
        value
            .as_str()
            .ok_or_else(|| Error::msg("kebab_case filter can only applied on strings"))?
            .to_case(Case::Kebab),
    ))
}

/// Change the case of a string into PascalCase.
fn pascal_case(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    Ok(Value::String(
        value
            .as_str()
            .ok_or_else(|| Error::msg("pascal_case filter can only applied on strings"))?
            .to_case(Case::Pascal),
    ))
}

/// Change the case of a string into snake_case.
fn snake_case(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    Ok(Value::String(
        value
            .as_str()
            .ok_or_else(|| Error::msg("snake_case filter can only applied on strings"))?
            .to_case(Case::Snake),
    ))
}

/// Pluralize form of a string.
fn pluralize(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    Ok(Value::String(to_plural(value.as_str().ok_or_else(
        || Error::msg("pluralize filter can only applied on strings"),
    )?)))
}

/// Singular form of a string.
fn singular(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    Ok(Value::String(to_singular(value.as_str().ok_or_else(
        || Error::msg("pluralize filter can only applied on strings"),
    )?)))
}

/// Change the case of an array of string into the case given by param `case`.
fn map_case(value: &Value, args: &HashMap<String, Value>) -> Result<Value> {
    let case = args
        .get("case")
        .ok_or_else(|| Error::msg("map_case filter requires a case argument"))?
        .as_str()
        .ok_or_else(|| Error::msg("map_case filter requires a case argument"))?;
    let value = value
        .as_array()
        .ok_or_else(|| Error::msg("map_case filter can only applied on array of strings"))?
        .iter()
        .map(|it| {
            it.as_str()
                .ok_or_else(|| Error::msg("map_case filter can only applied on array of strings"))
        })
        .collect::<Result<Vec<&str>>>()?;
    let case = match case {
        "camel" => Case::Camel,
        "kebab" => Case::Kebab,
        "pascal" => Case::Pascal,
        "snake" => Case::Snake,
        _ => return Err(Error::msg("map_case filter requires a case argument")),
    };
    Ok(Value::Array(
        value
            .into_iter()
            .map(|it| Value::String(it.to_case(case)))
            .collect(),
    ))
}

/// Add a suffix for an array of string.
fn map_suffix(value: &Value, args: &HashMap<String, Value>) -> Result<Value> {
    let suffix = args
        .get("suffix")
        .ok_or_else(|| Error::msg("map_suffix filter requires a suffix argument"))?
        .as_str()
        .ok_or_else(|| Error::msg("map_suffix filter requires a string suffix argument"))?;
    let value = value
        .as_array()
        .ok_or_else(|| Error::msg("map_suffix filter can only applied on array of strings"))?
        .iter()
        .map(|v| {
            v.as_str()
                .ok_or_else(|| Error::msg("map_suffix filter can only applied on array of strings"))
        })
        .collect::<Result<Vec<&str>>>()?;
    let value: Vec<Value> = value
        .into_iter()
        .map(|it| Value::String(format!("{}{}", it, suffix)))
        .collect();
    Ok(Value::Array(value))
}

/// Add a prefix for an array of string.
fn map_prefix(value: &Value, args: &HashMap<String, Value>) -> Result<Value> {
    let map_prefix = args
        .get("prefix")
        .ok_or_else(|| Error::msg("map_prefix filter requires a map_prefix argument"))?
        .as_str()
        .ok_or_else(|| Error::msg("map_prefix filter requires a string prefix argument"))?;
    let value = value
        .as_array()
        .ok_or_else(|| Error::msg("map_prefix filter can only applied on array of strings"))?
        .iter()
        .map(|v| {
            v.as_str()
                .ok_or_else(|| Error::msg("map_prefix filter can only applied on array of strings"))
        })
        .collect::<Result<Vec<&str>>>()?;
    let value: Vec<Value> = value
        .into_iter()
        .map(|it| Value::String(format!("{}{}", map_prefix, it)))
        .collect();
    Ok(Value::Array(value))
}

/// Register all filters above into the given tera template.
pub fn register(tera: &mut Tera) {
    tera.register_filter("camel_case", camel_case);
    tera.register_filter("kebab_case", kebab_case);
    tera.register_filter("pascal_case", pascal_case);
    tera.register_filter("snake_case", snake_case);
    tera.register_filter("pluralize", pluralize);
    tera.register_filter("singular", singular);
    tera.register_filter("map_prefix", map_prefix);
    tera.register_filter("map_suffix", map_suffix);
    tera.register_filter("map_case", map_case);
}
