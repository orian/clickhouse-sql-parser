
        CREATE TABLE jit_test_merge_tree (
            key UInt64,
            value_1 UInt64,
            value_2 UInt64,
            value_3 UInt64,
            value_4 UInt64,
            value_5 UInt64
        ) Engine = MergeTree
        ORDER BY key
    