
        CREATE TABLE minmax_wide
        (
            u128 UInt128, i128 Int128, small_i128 Int128, asc_i128 Int128,
            d128 Decimal128(4), d256 Decimal256(8),
            i256 Int256,
            cond UInt8,
            n_i128 Nullable(Int128)
        ) ENGINE = Memory
    