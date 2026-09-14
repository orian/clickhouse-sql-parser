-- Otorgar solo el permiso EXECUTE
GRANT ALTER TABLE EXECUTE ON my_iceberg_table TO my_user;

-- O conceder todos los permisos de ALTER TABLE (incluye ALTER TABLE EXECUTE)
GRANT ALTER TABLE ON my_iceberg_table TO my_user;
