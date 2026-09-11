SELECT number AS x, number + 1 AS y, 'Hello' AS z FROM numbers(10) FORMAT SQLInsert SETTINGS output_format_sql_insert_max_batch_size = 2
