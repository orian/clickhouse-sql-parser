clickhouse-cloud :) SELECT name FROM query_cache_test WHERE age > 1000 FORMAT Null SETTINGS use_query_cache=true;

SELECT name
FROM query_cache_test
WHERE age > 1000
FORMAT `Null`
SETTINGS use_query_cache = 1

Query id: a047527c-9d55-4e6b-9747-0ccad8787515

[c-red-qc-36-server-0] 2023.05.29 12:06:17.931007 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Debug\> executeQuery: (from 151.53.3.113:50412, user: tony) SELECT name FROM query_cache_test  WHERE age > 1000 FORMAT Null SETTINGS use_query_cache=true; (stage: Complete)
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931331 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Debug\> InterpreterSelectQuery: MergeTreeWhereOptimizer: condition "age > 1000" moved to PREWHERE
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931468 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Debug\> InterpreterSelectQuery: MergeTreeWhereOptimizer: condition "age > 1000" moved to PREWHERE
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931585 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Trace\> ContextAccess (tony): Access granted: SELECT(name, age) ON tony.test
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931696 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Trace\> ContextAccess (tony): Access granted: SELECT(name, age) ON tony.test
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931749 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Trace\> InterpreterSelectQuery: FetchColumns -> Complete
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931857 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Debug\> tony.test (e61a107c-e7f8-4445-825f-88f85c72f7e9) (SelectExecutor): Key condition: unknown
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931891 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Debug\> tony.test (e61a107c-e7f8-4445-825f-88f85c72f7e9) (SelectExecutor): Selected 1/1 parts by partition key, 1 parts by primary key, 12/12 marks by primary key, 12 marks to read from 1 ranges
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931913 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Trace\> tony.test (e61a107c-e7f8-4445-825f-88f85c72f7e9) (SelectExecutor): Spreading mark ranges among streams (default reading)
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931952 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Trace\> MergeTreeBaseSelectProcessor: PREWHERE condition was split into 1 steps: "greater(age, 1000)"
[c-red-qc-36-server-0] 2023.05.29 12:06:17.931975 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Trace\> MergeTreeInOrderSelectProcessor: Reading 1 ranges in order from part all_0_0_0, approx. 100000 rows starting from 0
[c-red-qc-36-server-0] 2023.05.29 12:06:17.932043 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Trace\> QueryCache: Entry found for query SELECT name FROM query_cache_test  WHERE age > 1000 FORMAT `Null` SETTINGS
[c-red-qc-36-server-0] 2023.05.29 12:06:17.932551 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Debug\> MemoryTracker: Peak memory usage (for query): 5.19 KiB.
[c-red-qc-36-server-0] 2023.05.29 12:06:17.932581 [ 454 ] {a047527c-9d55-4e6b-9747-0ccad8787515} \<Debug\> TCPHandler: Processed in 0.001961411 sec.
Ok.

0 rows in set. Elapsed: 0.002 sec.
