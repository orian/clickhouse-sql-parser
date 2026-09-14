CREATE TABLE table
(
    [...],
    vectors Array([Float64|Float32|BFloat16]),
    INDEX <index_name> vectors TYPE vector_similarity(<type>, <distance_function>, <dimensions>) [GRANULARITY <N>]
)
ENGINE = MergeTree
ORDER BY [...]
