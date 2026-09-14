CREATE TABLE [IF NOT EXISTS] [db.]table_name
[(
    name1 [type1],
    name2 [type2],
    ...
)]
ENGINE = BigQuery(project, dataset, table[, access_token][, key = value, ...])
