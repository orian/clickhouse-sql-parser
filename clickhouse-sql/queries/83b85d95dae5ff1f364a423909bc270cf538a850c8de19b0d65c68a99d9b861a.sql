
        CREATE TABLE jit_test_memory_nullable (
            key UInt64,
            value_1 Nullable(UInt64),
            value_2 Nullable(UInt64),
            value_3 Nullable(UInt64),
            value_4 Nullable(UInt64),
            value_5 Nullable(UInt64),
            predicate UInt8
        ) Engine = Memory
    