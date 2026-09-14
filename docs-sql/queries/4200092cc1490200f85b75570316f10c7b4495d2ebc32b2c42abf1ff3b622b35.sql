   CREATE MATERIALIZED VIEW schema.table_optimized
   REFRESH EVERY 1 HOUR
   ENGINE = ReplacingMergeTree()
   ORDER BY (user_id, event_date)
   AS SELECT * FROM schema.table_raw FINAL;
