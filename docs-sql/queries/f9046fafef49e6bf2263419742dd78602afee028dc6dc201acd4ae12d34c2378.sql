clickhouse-cloud :) SELECT 1 SETTINGS use_query_cache=true;

SELECT 1
SETTINGS use_query_cache = 1

Query id: a5a078c7-61e5-4036-a6f0-4d602d5b72d2

┌─1─┐
│ 1 │
└───┘

1 row in set. Elapsed: 0.001 sec.

clickhouse-cloud :) SELECT 1 SETTINGS use_query_cache=true;

SELECT 1
SETTINGS use_query_cache = 1

Query id: 322ae001-b1ab-463f-ac8d-dc5ba346f3f9

┌─1─┐
│ 1 │
└───┘

1 row in set. Elapsed: 0.001 sec.

clickhouse-cloud :) SELECT * FROM clusterAllReplicas(default,system.query_cache);

SELECT *
FROM clusterAllReplicas(default, system.query_cache)

Query id: c9b57eac-ba64-430e-8d51-8f865a13cc25

┌─query──────────────┬─result_size─┬─stale─┬─shared─┬─compressed─┬──────────expires_at─┬─────────────key_hash─┐
│ SELECT 1 SETTINGS  │         136 │     0 │      1 │          1 │ 2023-08-02 15:08:23 │ 12188185624808016954 │
└────────────────────┴─────────────┴───────┴────────┴────────────┴─────────────────────┴──────────────────────┘

1 row in set. Elapsed: 0.005 sec.

clickhouse-cloud :) SELECT * FROM clusterAllReplicas(default,system.events) WHERE event LIKE 'QueryCache%'

SELECT *
FROM clusterAllReplicas(default, system.events)
WHERE event LIKE 'QueryCache%'

Query id: d536555e-b8ab-4cd4-9741-c04e95612bec

┌─event────────────┬─value─┬─description────────────────────────────────────────────────────────────────────────────────────────────┐
│ QueryCacheHits   │     1 │ Number of times a query result has been found in the query cache (and query computation was avoided).  │
│ QueryCacheMisses │     1 │ Number of times a query result has not been found in the query cache (and required query computation). │
└──────────────────┴───────┴────────────────────────────────────────────────────────────────────────────────────────────────────────┘

2 rows in set. Elapsed: 0.006 sec.
