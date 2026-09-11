CREATE TABLE tab
(
    a Int64 STATISTICS(tdigest, uniq),
    b Float64
)
ENGINE = MergeTree
ORDER BY a
