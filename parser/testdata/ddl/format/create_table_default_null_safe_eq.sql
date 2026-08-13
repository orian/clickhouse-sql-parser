-- Origin SQL:
CREATE TABLE db.things (`name` String, `matches` Bool DEFAULT other_name <=> name, `other_name` Nullable(String)) ENGINE = MergeTree ORDER BY name;


-- Format SQL:
CREATE TABLE db.things (`name` String, `matches` Bool DEFAULT other_name <=> name, `other_name` Nullable(String)) ENGINE = MergeTree ORDER BY name;
