SET output_format_sql_insert_table_name = 'some_table';
SELECT * FROM some_data
INTO OUTFILE 'dump.sql'
FORMAT SQLInsert
