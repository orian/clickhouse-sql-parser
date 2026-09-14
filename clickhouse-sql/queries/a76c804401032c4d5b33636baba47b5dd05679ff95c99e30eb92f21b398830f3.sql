
        CREATE TABLE sort_wide_16 (
            k1 UInt64, k2 UInt64,
            p01 UInt64, p02 UInt64, p03 UInt64, p04 UInt64,
            p05 UInt64, p06 UInt64, p07 UInt64, p08 UInt64,
            p09 UInt64, p10 UInt64, p11 UInt64, p12 UInt64,
            p13 UInt64, p14 UInt64, p15 UInt64, p16 UInt64
        ) ENGINE = MergeTree ORDER BY tuple()
    