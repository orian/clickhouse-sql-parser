-- uniqCombined64 with large dataset
SELECT uniqCombined64(number) FROM numbers(1e10);

-- uniqCombined with same dataset shows poor approximation
SELECT uniqCombined(number) FROM numbers(1e10);
