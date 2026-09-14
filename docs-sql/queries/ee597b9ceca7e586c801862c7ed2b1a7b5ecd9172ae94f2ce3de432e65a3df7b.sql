EXPLAIN indexes = 1
SELECT count() AS c
FROM logs
WHERE code = '200'

┌─explain────────────────────────────────────────────────────────────┐
│ Expression ((Project names + Projection))                          │
│   AggregatingProjection                                            │
│     Expression (Before GROUP BY)                                   │
│       Filter ((WHERE + Change column names to column identifiers)) │
│         ReadFromMergeTree (default.logs)                           │
│         Indexes:                                                   │
│           PrimaryKey                                               │
│             Keys:                                                  │
│               code                                                 │
│             Condition: (code in ['200', '200'])                    │
│             Parts: 3/3 │
│             Granules: 8012/12209 │
│     ReadFromPreparedSource (_minmax_count_projection)              │
└────────────────────────────────────────────────────────────────────┘
