WITH values AS (
  SELECT map('ClickHouse', 3) AS value
  UNION ALL
  SELECT map('ClickBench', 2, 'ClickHouse', 4) AS value
)
SELECT sumMap(value)
FROM values;
