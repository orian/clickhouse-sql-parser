CREATE TABLE test (i UInt32, s String) ENGINE MergeTree ORDER BY i;
SELECT constraint_catalog,
       constraint_schema,
       constraint_name,
       table_catalog,
       table_schema,
       table_name,
       column_name,
       ordinal_position,
       position_in_unique_constraint,
       referenced_table_schema,
       referenced_table_name,
       referenced_column_name
FROM information_schema.key_column_usage 
WHERE table_name = 'test' 
FORMAT Vertical;
