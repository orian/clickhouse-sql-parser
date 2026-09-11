SET enable_quantized_codec = 1;

CREATE TABLE vectors
(
    id UInt32,
    vec Array(BFloat16) CODEC(Quantized('rabitq', 1536))
)
ENGINE = MergeTree
ORDER BY ...;
