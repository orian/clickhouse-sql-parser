SELECT referenced_data_file, row_number
FROM file(deletes.puffin, Puffin)
ARRAY JOIN deleted_rows AS row_number
ORDER BY referenced_data_file, row_number;
