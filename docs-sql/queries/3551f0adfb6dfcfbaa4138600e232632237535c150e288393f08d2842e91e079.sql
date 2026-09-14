EXPLAIN PIPELINE
SELECT *
FROM logs
WHERE (code = '200') AND (timestamp >= '2025-01-01 00:00:00') AND (timestamp <= '2026-01-01 00:00:00')
ORDER BY timestamp ASC
LIMIT 10

┌─explain───────────────────────────────────────────────────────────────────────┐
│ (Expression)                                                                  │
│ ExpressionTransform                                                           │
│   (Limit)                                                                     │
│   Limit │
│     (Sorting)                                                                 │
│     MergingSortedTransform 12 → 1 │
│       MergeSortingTransform × 12 │
│         LimitsCheckingTransform × 12 │
│           PartialSortingTransform × 12 │
│             (Expression)                                                      │
│             ExpressionTransform × 12 │
│               (Expression)                                                    │
│               ExpressionTransform × 12 │
│                 (ReadFromMergeTree)                                           │
│                 MergeTreeSelect(pool: ReadPool, algorithm: Thread) × 12 0 → 1 │
└───────────────────────────────────────────────────────────────────────────────┘

15 rows in set. Elapsed: 0.004 sec.
