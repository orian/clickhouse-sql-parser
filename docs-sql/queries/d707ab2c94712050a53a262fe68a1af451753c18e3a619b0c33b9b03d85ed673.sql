CREATE TABLE db1.source_table
(
    city VARCHAR,
    country VARCHAR,
    continent VARCHAR
)
engine = MergeTree()
ORDER BY continent;
