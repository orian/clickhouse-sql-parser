-- Origin SQL:
GRANT SELECT(x,y) ON db.table TO john;
GRANT SELECT(x,y) ON db.table TO john WITH GRANT OPTION WITH REPLACE OPTION;
GRANT SELECT ON db.* TO john;
GRANT SELECT(x,y) ON table TO john;
GRANT SELECT ON *.* TO john;
GRANT SELECT(x,y) ON db.table TO CURRENT_USER;
GRANT SELECT(x,y) ON db.table TO CURRENT_USER,john,mary;
GRANT ALL ON *.* TO admin_role WITH GRANT OPTION;
GRANT SELECT,INSERT ON database.table_1 TO table_1_select_role;
GRANT SELECT(x, y, z),INSERT ON database.table_1 TO table_1_select_role;
GRANT SELECT, dictGet ON *.*  TO select_all_role;
GRANT SELECT ON *.* TO select_all_role WITH REPLACE OPTION;


-- Beautify SQL:
GRANT SELECT(x, y) ON db.table TO john;
GRANT SELECT(x, y) ON db.table TO john WITH GRANT OPTION WITH REPLACE OPTION;
GRANT SELECT ON db.* TO john;
GRANT SELECT(x, y) ON table TO john;
GRANT SELECT ON *.* TO john;
GRANT SELECT(x, y) ON db.table TO CURRENT_USER;
GRANT SELECT(x, y) ON db.table TO CURRENT_USER, john, mary;
GRANT ALL ON *.* TO admin_role WITH GRANT OPTION;
GRANT SELECT, INSERT ON database.table_1 TO table_1_select_role;
GRANT SELECT(x, y, z), INSERT ON database.table_1 TO table_1_select_role;
GRANT SELECT, dictGet ON *.* TO select_all_role;
GRANT SELECT ON *.* TO select_all_role WITH REPLACE OPTION;
