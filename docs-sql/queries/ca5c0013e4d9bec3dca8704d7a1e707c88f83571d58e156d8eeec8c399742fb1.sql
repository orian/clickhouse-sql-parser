try=# EXPLAIN (VERBOSE, COSTS OFF)
       SELECT start_at, duration, resource FROM logs WHERE req_id = 4117909262;
                                             QUERY PLAN
-----------------------------------------------------------------------------------------------------
 Foreign Scan on public.logs
   Output: start_at, duration, resource
   Remote SQL: SELECT start_at, duration, resource FROM "default".logs WHERE ((req_id = 4117909262))
(3 rows)
