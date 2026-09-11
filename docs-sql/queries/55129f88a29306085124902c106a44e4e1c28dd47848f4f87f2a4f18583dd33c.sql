    CREATE ROW POLICY A_row_filter ON db1.table1 FOR SELECT USING column1 = 'A' TO A_rows_users;
