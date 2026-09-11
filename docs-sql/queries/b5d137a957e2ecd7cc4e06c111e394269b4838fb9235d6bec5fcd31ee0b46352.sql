-- Conceder apenas a permissão EXECUTE
GRANT ALTER TABLE EXECUTE ON my_iceberg_table TO my_user;

-- Ou conceder todas as permissões de ALTER TABLE (inclui ALTER TABLE EXECUTE)
GRANT ALTER TABLE ON my_iceberg_table TO my_user;
