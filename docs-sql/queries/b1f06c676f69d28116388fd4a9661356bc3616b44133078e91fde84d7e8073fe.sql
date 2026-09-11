SELECT parseDateTime64BestEffort('2025-01-01') AS a, toTypeName(a) AS t
UNION ALL
SELECT parseDateTime64BestEffort('2025-01-01 01:01:00.12346') AS a, toTypeName(a) AS t
UNION ALL
SELECT parseDateTime64BestEffort('2025-01-01 01:01:00.12346',6) AS a, toTypeName(a) AS t
UNION ALL
SELECT parseDateTime64BestEffort('2025-01-01 01:01:00.12346',3,'Asia/Istanbul') AS a, toTypeName(a) AS t
FORMAT PrettyCompactMonoBlock
