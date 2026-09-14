CREATE TABLE perf_lc_fixed_str_groupby(
            a LowCardinality(FixedString(14)),
            b LowCardinality(FixedString(14))
        ) ENGINE MergeTree ORDER BY tuple()
    