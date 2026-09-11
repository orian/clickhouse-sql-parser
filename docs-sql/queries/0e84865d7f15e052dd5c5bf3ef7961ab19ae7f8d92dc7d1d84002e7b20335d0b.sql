EXPLAIN indexes=1
SELECT *
FROM events_rollup_1h
WHERE bucket_start BETWEEN now() - INTERVAL 3 DAY AND now()
  AND country = 'US';
