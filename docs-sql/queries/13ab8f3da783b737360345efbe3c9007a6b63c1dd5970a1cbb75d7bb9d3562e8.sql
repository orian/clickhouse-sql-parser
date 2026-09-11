CREATE TABLE adaptive
(
    time DateTime,
    user_id UInt64
)
ENGINE = MergeTree
ORDER BY time
SETTINGS allow_experimental_adaptive_codec_selection = 1;

INSERT INTO adaptive SELECT toDateTime('2026-01-01') + number, cityHash64(number) FROM numbers(1000000);
OPTIMIZE TABLE adaptive FINAL;
