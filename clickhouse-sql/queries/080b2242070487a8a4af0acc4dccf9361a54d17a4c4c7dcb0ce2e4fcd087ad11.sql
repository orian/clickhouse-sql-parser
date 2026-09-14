CREATE TABLE perf_avg(
        num UInt64,
        num_u Decimal256(75) MATERIALIZED toDecimal256(num / 400000, 75),
        num_f Float64 MATERIALIZED num / 100
        ) ENGINE = MergeTree() ORDER BY num
    