CREATE TABLE sometable
(
    `path` String,
    `month` Date,
    `hits` UInt32
)
ENGINE = MergeTree
ORDER BY tuple(month, path)
