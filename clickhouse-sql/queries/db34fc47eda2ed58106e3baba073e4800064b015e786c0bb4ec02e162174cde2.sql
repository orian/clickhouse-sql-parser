
    SELECT
      multiIf(a >= 100000 AND a < 10000000, 0, 1) AS index,
      COUNT(DISTINCT b) AS total
    FROM test
    GROUP BY 1
    SETTINGS max_threads = {threads};
  