 EXPLAIN select count(*) from taxi.trips;
                    QUERY PLAN
 -------------------------------------------------
  Foreign Scan  (cost=1.00..-0.90 rows=1 width=8)
    Relations: Aggregate on (trips)
 (2 rows)
