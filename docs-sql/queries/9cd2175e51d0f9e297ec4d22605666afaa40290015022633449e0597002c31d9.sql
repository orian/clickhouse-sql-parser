CREATE TABLE hits_clickhouse
(
    url String,
    eventtime DateTime,
    counterid UInt32
)
ENGINE = MergeTree()
ORDER BY (counterid, eventtime);
