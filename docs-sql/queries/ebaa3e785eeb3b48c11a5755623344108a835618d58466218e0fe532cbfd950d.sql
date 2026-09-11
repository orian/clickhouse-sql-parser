CREATE TABLE default.sample2
(
    `x` Int8,
    `y` Int8 DEFAULT 0
)
ENGINE = MergeTree
ORDER BY x
