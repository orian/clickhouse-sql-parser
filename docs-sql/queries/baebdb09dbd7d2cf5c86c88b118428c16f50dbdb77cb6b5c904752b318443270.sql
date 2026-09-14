EXPLAIN PIPELINE
SELECT *
FROM logs
WHERE (code = '200') AND (timestamp >= '2025-01-01 00:00:00') AND (timestamp <= '2026-01-01 00:00:00')
ORDER BY toUnixTimestamp(timestamp) ASC
LIMIT 10

┌─explain──────────────────────────────────────────────────────────────────────────┐
│ (Expression)                                                                     │
│ ExpressionTransform                                                              │
│   (Limit)                                                                        │
│   Limit │
│     (Sorting)                                                                    │
│     MergingSortedTransform 3 → 1 │
│       BufferChunks × 3 │
│         (Expression)                                                             │
│         ExpressionTransform × 3 │
│           (Expression)                                                           │
│           ExpressionTransform × 3 │
│             (ReadFromMergeTree)                                                  │
│             MergeTreeSelect(pool: ReadPoolInOrder, algorithm: InOrder) × 3 0 → 1 │
└──────────────────────────────────────────────────────────────────────────────────┘

13 rows in set. Elapsed: 0.003 sec.
