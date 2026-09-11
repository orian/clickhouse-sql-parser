try=# EXPLAIN (VERBOSE, COSTS OFF)
       SELECT count(*), sum(val), min(ts), max(ts) FROM events;
                                                       QUERY PLAN
-------------------------------------------------------------------------------------------------------------------------
 Finalize Aggregate
   Output: count(*), sum(events.val), min(events.ts), max(events.ts)
   ->  Append
         ->  Foreign Scan
               Output: (PARTIAL count(*)), (PARTIAL sum(events.val)), (PARTIAL min(events.ts)), (PARTIAL max(events.ts))
               Relations: Aggregate on (events_2023 events)
               Remote SQL: SELECT count(*), sum(val), min(ts), max(ts) FROM "default".events
         ->  Partial Aggregate
               Output: PARTIAL count(*), PARTIAL sum(events_1.val), PARTIAL min(events_1.ts), PARTIAL max(events_1.ts)
               ->  Seq Scan on public.events_2024 events_1
                     Output: events_1.val, events_1.ts
