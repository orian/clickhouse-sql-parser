SELECT clickhouse_raw_query(
    'SELECT schema_name, schema_owner from information_schema.schemata',
    'host=localhost port=8123'
);
