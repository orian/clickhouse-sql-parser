SELECT * FROM db.my_tables -- 許可済み
SELECT * FROM db.my_tables_0 -- 許可済み
SELECT * FROM db.my_tables_1 -- 許可済み

SELECT * FROM db.other_table -- 未許可
SELECT * FROM db2.my_tables -- 未許可
