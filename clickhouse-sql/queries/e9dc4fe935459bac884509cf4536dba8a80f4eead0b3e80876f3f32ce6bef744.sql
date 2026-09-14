
        CREATE TABLE IF NOT EXISTS source_table
        (
        id UInt32,
        col_int8 Nullable(Int8),
        col_uint8 Nullable(UInt8),
        col_int16 Nullable(Int16),
        col_uint16 Nullable(UInt16),
        col_int32 Nullable(Int32),
        col_uint32 Nullable(UInt32),
        col_int64 Nullable(Int64),
        col_uint64 Nullable(UInt64),
        col_float32 Nullable(Float32),
        col_float64 Nullable(Float64),
        col_string Nullable(String),
        col_fixed_string Nullable(FixedString(10)),
        col_date Nullable(Date),
        col_datetime Nullable(DateTime),
        col_datetime64 Nullable(DateTime64(3)),
        col_decimal32 Nullable(Decimal32(5)),
        col_decimal64 Nullable(Decimal64(10)),
        col_decimal128 Nullable(Decimal128(20)),
        col_array Array(Nullable(UInt32)),
        col_low_cardinality LowCardinality(Nullable(String))
        ) ENGINE = MergeTree()
        ORDER BY id
    