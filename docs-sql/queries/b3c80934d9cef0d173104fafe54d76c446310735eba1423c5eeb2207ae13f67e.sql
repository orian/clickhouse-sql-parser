CREATE TABLE default.sample
(
    `x` Int8,
    `y` Nullable(Int8)
)
ENGINE = MergeTree
ORDER BY x
