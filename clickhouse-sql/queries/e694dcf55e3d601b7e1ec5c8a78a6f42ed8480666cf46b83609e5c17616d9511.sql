
        CREATE TABLE t_uniq_stats(c1 UInt64 STATISTICS(uniq), c2 LowCardinality(String) STATISTICS(uniq), c3 LowCardinality(String) STATISTICS(uniq), c4 LowCardinality(String) STATISTICS(uniq), c5 LowCardinality(String) STATISTICS(uniq))
        ENGINE = MergeTree ORDER BY c1;
    