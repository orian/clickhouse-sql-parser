CREATE TABLE tab
(
    id Int32,
    vec Array(BFloat16) CODEC(NONE),
    INDEX idx vec TYPE vector_similarity('hnsw', 'L2Distance', 2)
)
ENGINE = MergeTree ORDER BY id;
