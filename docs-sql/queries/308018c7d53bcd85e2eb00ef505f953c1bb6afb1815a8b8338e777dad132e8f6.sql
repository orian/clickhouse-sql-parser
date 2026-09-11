SET decimal_check_overflow = 0;
SELECT toDecimal32(4.2, 8) AS x, 6 * x
