CREATE TABLE db1.table1 (
   id UInt64,
   column1 String,
   column2 String
)
ENGINE MergeTree
ORDER BY id;
