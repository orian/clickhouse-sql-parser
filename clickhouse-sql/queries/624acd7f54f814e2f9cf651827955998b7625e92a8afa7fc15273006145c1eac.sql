
    SELECT number
    FROM numbers_mt(1e8) AS n any inner join
    (
        SELECT 157::UInt64 AS p
        FROM system.one
    ) AS o
    on number = p
    FORMAT Null SETTINGS max_threads=1, join_algorithm='hash'
  