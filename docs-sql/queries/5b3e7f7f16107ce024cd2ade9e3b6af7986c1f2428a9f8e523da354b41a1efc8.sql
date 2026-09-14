CREATE TABLE coalescing_tuples
(
    key UInt64,
    data Tuple(
        value_a Nullable(UInt64),
        value_b Nullable(String),
        nested Tuple(
            value_c Nullable(UInt64)
        )
    )
) ENGINE = CoalescingMergeTree()
ORDER BY key
SETTINGS allow_tuple_element_aggregation = 1;

INSERT INTO coalescing_tuples VALUES (1, (100, NULL, (NULL)));
INSERT INTO coalescing_tuples VALUES (1, (NULL, 'hello', (42)));

SELECT key, data.value_a, data.value_b, data.nested.value_c FROM coalescing_tuples FINAL;
