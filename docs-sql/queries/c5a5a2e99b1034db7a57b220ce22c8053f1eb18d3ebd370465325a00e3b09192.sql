CREATE DICTIONARY hierarchical_dictionary
(
    id UInt64,
    parent_id UInt64 HIERARCHICAL,
    name String
)
PRIMARY KEY id
SOURCE(CLICKHOUSE(HOST 'localhost' PORT 9000 USER 'default' TABLE 'hierarchy_source' DB 'default'))
LAYOUT(HASHED())
LIFETIME(MIN 300 MAX 600);
