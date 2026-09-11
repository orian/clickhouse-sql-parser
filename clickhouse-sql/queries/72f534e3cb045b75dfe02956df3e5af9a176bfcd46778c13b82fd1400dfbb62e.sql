
        CREATE TABLE jit_test_merge_tree_nullable (
            key UInt64,
            value_1 Nullable(UInt64),
            value_2 Nullable(UInt64),
            value_3 Nullable(UInt64),
            value_4 Nullable(UInt64),
            value_5 Nullable(UInt64)
        ) Engine = MergeTree
        ORDER BY key
    