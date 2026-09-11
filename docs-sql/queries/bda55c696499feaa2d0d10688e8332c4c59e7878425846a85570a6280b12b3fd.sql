CREATE HYPOTHETICAL PROJECTION [IF NOT EXISTS] name
    ON [db.]table_name (SELECT <columns> [WHERE ...] [GROUP BY ...] [ORDER BY ...]) [WITH SETTINGS (...)]

CREATE HYPOTHETICAL PROJECTION [IF NOT EXISTS] name
    ON [db.]table_name INDEX <expression> TYPE <projection_index_type> [WITH SETTINGS (...)]
