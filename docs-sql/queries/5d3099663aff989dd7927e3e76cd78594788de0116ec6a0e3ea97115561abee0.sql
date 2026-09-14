SELECT count()
FROM logs
WHERE (code = '200') AND (timestamp >= '2025-01-01 00:00:00') AND (timestamp <= '2026-01-01 00:00:00')

┌─count()─┐
│  689742 │
└─────────┘

1 row in set. Elapsed: 0.008 sec. Processed 712.70 thousand rows, 6.41 MB (88.92 million rows/s., 799.27 MB/s.)

EXPLAIN indexes = 1
SELECT count()
FROM logs
WHERE (code = '200') AND (timestamp >= '2025-01-01 00:00:00') AND (timestamp <= '2026-01-01 00:00:00')

┌─explain───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ Expression ((Project names + Projection))                                                                                                                         │
│   Aggregating                                                                                                                                                     │
│     Expression (Before GROUP BY)                                                                                                                                  │
│       Expression                                                                                                                                                  │
│         ReadFromMergeTree (default.logs)                                                                                                                          │
│         Indexes:                                                                                                                                                  │
│           PrimaryKey                                                                                                                                              │
│             Keys:                                                                                                                                                 │
│               code                                                                                                                                                │
│               toUnixTimestamp(timestamp)                                                                                                                          │
│             Condition: and((toUnixTimestamp(timestamp) in (-Inf, 1767225600]), and((toUnixTimestamp(timestamp) in [1735689600, +Inf)), (code in ['200', '200']))) │
│             Parts: 3/3 │
│             Granules: 87/12209 │
└───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

13 rows in set. Elapsed: 0.002 sec.

