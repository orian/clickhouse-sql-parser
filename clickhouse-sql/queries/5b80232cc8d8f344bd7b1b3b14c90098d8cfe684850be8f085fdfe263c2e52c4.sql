
        CREATE TABLE test_table_medium_{array_type}
        (
            `set` Array({array_type}),
            `subset` Array ({array_type})
        )
        ENGINE = MergeTree ORDER BY set;
    