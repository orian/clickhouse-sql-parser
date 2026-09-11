create table default.my_first_table
(
    user_id String,
    json String,
    name String,
) ENGINE = MergeTree()
ORDER BY user_id;
