-- EXECUTE権限のみを付与する
GRANT ALTER TABLE EXECUTE ON my_iceberg_table TO my_user;

-- またはすべてのALTER TABLE権限を付与する（ALTER TABLE EXECUTEを含む）
GRANT ALTER TABLE ON my_iceberg_table TO my_user;
