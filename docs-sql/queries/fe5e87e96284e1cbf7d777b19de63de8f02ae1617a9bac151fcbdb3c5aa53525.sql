CREATE TABLE [IF NOT EXISTS] [db.]table_name
(
    name1 [type1],
    name2 [type2],
    ...
) ENGINE = Remote(addresses_expr, [db, table, [user [, password], sharding_key]])
[SETTINGS name = value, ...]
