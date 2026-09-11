
        CREATE TABLE t_nullable
        (
        key_string Nullable(String),
        key_string_two_level Nullable(String),
        key_fixed_string Nullable(FixedString(3)),
        key_int64 Nullable(Int64),
        key_int64_two_level Nullable(Int64),
        key_int32 Nullable(Int32),
        key_int32_two_level Nullable(Int32),
        Key_int16 Nullable(Int16),
        key_int8 Nullable(Int8),
        i1 Nullable(Int64),
        i2 Nullable(Int64),
        )
        ENGINE = MergeTree
        ORDER BY tuple()
    