SELECT * FROM db.my_tables -- 권한 부여됨
SELECT * FROM db.my_tables_0 -- 권한 부여됨
SELECT * FROM db.my_tables_1 -- 권한 부여됨

SELECT * FROM db.other_table -- 권한 부여 안 됨
SELECT * FROM db2.my_tables -- 권한 부여 안 됨
