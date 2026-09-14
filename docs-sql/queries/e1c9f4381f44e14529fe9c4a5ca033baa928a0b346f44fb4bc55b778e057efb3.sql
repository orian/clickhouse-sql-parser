    CREATE ROW POLICY allow_other_users_filter 
    ON db1.table1 FOR SELECT USING 1 TO clickhouse_admin, column1_users;
