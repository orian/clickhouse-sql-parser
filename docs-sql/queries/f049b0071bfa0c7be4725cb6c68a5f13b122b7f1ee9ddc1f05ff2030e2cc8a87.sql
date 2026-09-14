-- EXECUTE 권한만 부여
GRANT ALTER TABLE EXECUTE ON my_iceberg_table TO my_user;

-- 또는 모든 ALTER TABLE 권한 부여 (ALTER TABLE EXECUTE 포함)
GRANT ALTER TABLE ON my_iceberg_table TO my_user;
