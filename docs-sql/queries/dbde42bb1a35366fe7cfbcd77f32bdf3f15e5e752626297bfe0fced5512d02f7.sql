-- Single-policy rename
ALTER ROW POLICY p1 ON db.table RENAME TO p1_new;

-- Rename plus other alterations on the same single policy
ALTER POLICY old_name ON db.table RENAME TO new_name USING id > 10;

-- Multiple names, one table
ALTER POLICY p1, p2 ON db.table TO ALL;

-- One name, multiple tables
ALTER POLICY p1 ON db.table, db.table2 USING NONE;

-- Mixed targets without rename
ALTER POLICY p1 ON db.table, p2 ON db2.table2 TO ALL;
