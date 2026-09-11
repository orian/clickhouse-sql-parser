-- 仅授予 EXECUTE 权限
GRANT ALTER TABLE EXECUTE ON my_iceberg_table TO my_user;

-- 或授予所有 ALTER TABLE 权限（包含 ALTER TABLE EXECUTE）
GRANT ALTER TABLE ON my_iceberg_table TO my_user;
