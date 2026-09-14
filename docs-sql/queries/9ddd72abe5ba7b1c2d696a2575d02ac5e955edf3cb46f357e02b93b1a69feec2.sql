CREATE TABLE db1.target_table_population
(
    `city` String,
    `country` String,
    `continent` String,
    `population` UInt16,
)
ENGINE = MergeTree
ORDER BY continent;
