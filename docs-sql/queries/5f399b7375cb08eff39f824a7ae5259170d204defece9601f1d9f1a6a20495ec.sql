SELECT
    avg(y),
    avg(coalesce(y, 0))
FROM t_null_big
