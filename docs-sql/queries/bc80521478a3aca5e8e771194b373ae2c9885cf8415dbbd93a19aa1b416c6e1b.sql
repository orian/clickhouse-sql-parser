SELECT * FROM db.my_tables -- выдан
SELECT * FROM db.my_tables_0 -- выдан
SELECT * FROM db.my_tables_1 -- выдан

SELECT * FROM db.other_table -- не_выдан
SELECT * FROM db2.my_tables -- не_выдан
