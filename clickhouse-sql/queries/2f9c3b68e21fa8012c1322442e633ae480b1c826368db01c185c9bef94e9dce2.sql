
        CREATE TABLE na_src
        (
            i64    Int64,
            n_i64  Nullable(Int64),
            n_u32  Nullable(UInt32),
            b      Bool,
            d      Date,
            dt64   DateTime64(3),
            dec    Decimal64(4),
            s      String,
            lc     LowCardinality(String),
            arr    Array(Int64)
        ) ENGINE = Memory
    