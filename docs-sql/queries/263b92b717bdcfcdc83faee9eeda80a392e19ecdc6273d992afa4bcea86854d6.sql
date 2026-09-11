   -- When adding new tables to the ClickPipe, you'll need to add them to the publication manually as well.
   CREATE PUBLICATION clickpipes FOR TABLE table_to_replicate, table_to_replicate2, TABLES IN SCHEMA _timescaledb_internal;
