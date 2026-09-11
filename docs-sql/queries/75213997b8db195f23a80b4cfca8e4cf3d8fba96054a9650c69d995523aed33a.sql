SELECT *
FROM system.server_settings
WHERE name IN ('max_table_num_to_warn', 'max_table_num_to_throw', 'max_database_num_to_warn', 'max_database_num_to_throw', 'max_view_num_to_warn', 'max_view_num_to_throw', 'max_dictionary_num_to_warn', 'max_dictionary_num_to_throw', 'max_named_collection_num_to_warn', 'max_named_collection_num_to_throw')
