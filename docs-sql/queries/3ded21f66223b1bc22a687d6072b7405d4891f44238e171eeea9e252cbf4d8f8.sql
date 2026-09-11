-- This example shows that `count(DISTINCT num)` is performed by the `uniqExact` function according to the `count_distinct_implementation` setting value.
SELECT name, value FROM system.settings WHERE name = 'count_distinct_implementation';
SELECT count(DISTINCT num) FROM t
