CREATE TABLE perf_lc_str(
        str LowCardinality(String),
        arr Array(LowCardinality(String)) default [str]
        ) ENGINE = MergeTree() ORDER BY tuple()
    