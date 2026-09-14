CREATE TABLE perf_lc_num(
        num UInt8,
        arr Array(LowCardinality(Int64)) default [num]
        ) ENGINE = StripeLog
    