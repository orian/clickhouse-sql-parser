
        CREATE TABLE t_basic_stats(c1 UInt64 STATISTICS(basic), c2 UInt64 STATISTICS(basic), c3 UInt64 STATISTICS(basic), c4 UInt64 STATISTICS(basic), c5 UInt64 STATISTICS(basic), c6 UInt64 STATISTICS(basic), c7 UInt64 STATISTICS(basic), c8 UInt64 STATISTICS(basic), c9 UInt64 STATISTICS(basic), c10 UInt64 STATISTICS(basic))
        ENGINE = MergeTree ORDER BY c1;
    