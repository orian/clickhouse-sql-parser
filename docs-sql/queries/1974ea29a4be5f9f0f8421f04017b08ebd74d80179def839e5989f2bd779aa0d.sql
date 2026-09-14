EXPLAIN ANALYZE SELECT number % 10 AS k, count() FROM numbers_mt(1000000) GROUP BY k;
