   SELECT database, table, mutation_id, command, create_time, is_done
   FROM system.mutations
   WHERE NOT is_done
   ORDER BY create_time DESC;
