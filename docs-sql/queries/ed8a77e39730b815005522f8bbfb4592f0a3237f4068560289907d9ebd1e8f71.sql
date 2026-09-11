clickhouse-cloud :) SELECT name FROM query_cache_test WHERE age > 1000 FORMAT Null SETTINGS use_query_cache=true;

SELECT name
FROM query_cache_test
WHERE age > 1000
FORMAT `Null`
SETTINGS use_query_cache = 1

Query id: 3754a7fd-b786-47c1-a258-dfbc75e35a04

[c-red-qc-36-server-0] 2023.05.29 12:06:10.542408 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> executeQuery: (from 151.53.3.113:50412, user: tony) SELECT name FROM query_cache_test  WHERE age > 1000 FORMAT Null SETTINGS use_query_cache=true; (stage: Complete)
[c-red-qc-36-server-0] 2023.05.29 12:06:10.542744 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> InterpreterSelectQuery: MergeTreeWhereOptimizer: condition "age > 1000" moved to PREWHERE
[c-red-qc-36-server-0] 2023.05.29 12:06:10.542900 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> InterpreterSelectQuery: MergeTreeWhereOptimizer: condition "age > 1000" moved to PREWHERE
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543020 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> ContextAccess (tony): Access granted: SELECT(name, age) ON tony.test
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543164 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> ContextAccess (tony): Access granted: SELECT(name, age) ON tony.test
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543226 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> InterpreterSelectQuery: FetchColumns -> Complete
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543337 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> tony.test (e61a107c-e7f8-4445-825f-88f85c72f7e9) (SelectExecutor): Key condition: unknown
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543395 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> tony.test (e61a107c-e7f8-4445-825f-88f85c72f7e9) (SelectExecutor): Selected 1/1 parts by partition key, 1 parts by primary key, 12/12 marks by primary key, 12 marks to read from 1 ranges
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543412 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> tony.test (e61a107c-e7f8-4445-825f-88f85c72f7e9) (SelectExecutor): Spreading mark ranges among streams (default reading)
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543461 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> MergeTreeBaseSelectProcessor: PREWHERE condition was split into 1 steps: "greater(age, 1000)"
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543484 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> MergeTreeInOrderSelectProcessor: Reading 1 ranges in order from part all_0_0_0, approx. 100000 rows starting from 0
[c-red-qc-36-server-0] 2023.05.29 12:06:10.543559 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> QueryCache: No entry found for query SELECT name FROM query_cache_test  WHERE age > 1000 FORMAT `Null` SETTINGS
[c-red-qc-36-server-0] 2023.05.29 12:06:10.547760 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Trace\> QueryCache: Stored result of query SELECT name FROM query_cache_test  WHERE age > 1000 FORMAT `Null` SETTINGS
[c-red-qc-36-server-0] 2023.05.29 12:06:10.547827 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> executeQuery: Read 100000 rows, 97.66 KiB in 0.005508 sec., 18155410.31227306 rows/sec., 17.31 MiB/sec.
[c-red-qc-36-server-0] 2023.05.29 12:06:10.547913 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> MemoryTracker: Peak memory usage (for query): 451.89 KiB.
[c-red-qc-36-server-0] 2023.05.29 12:06:10.547933 [ 454 ] {3754a7fd-b786-47c1-a258-dfbc75e35a04} \<Debug\> TCPHandler: Processed in 0.005911032 sec.
Ok.

0 rows in set. Elapsed: 0.006 sec. Processed 100.00 thousand rows, 100.00 KB (17.56 million rows/s., 17.56 MB/s.)
