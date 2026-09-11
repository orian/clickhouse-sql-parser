-- Предоставить только разрешение EXECUTE
GRANT ALTER TABLE EXECUTE ON my_iceberg_table TO my_user;

-- Или предоставить все разрешения ALTER TABLE (включает ALTER TABLE EXECUTE)
GRANT ALTER TABLE ON my_iceberg_table TO my_user;
