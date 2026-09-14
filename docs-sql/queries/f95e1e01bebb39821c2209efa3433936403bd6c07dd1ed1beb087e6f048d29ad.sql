CREATE TABLE users_a (
    uid Int16,
    name String,
    age Int16,
    name_len UInt8 MATERIALIZED length(name),
    CONSTRAINT c1 ASSUME length(name) = name_len
)
ENGINE=MergeTree
ORDER BY (name_len, name);
