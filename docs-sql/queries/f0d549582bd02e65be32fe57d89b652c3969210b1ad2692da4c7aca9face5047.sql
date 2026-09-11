 CREATE OR REPLACE TABLE test
 (
     id UInt64,
     Name TEXT,
     updatedAt DateTime DEFAULT now()
 )
 ENGINE = MergeTree
 ORDER BY id;
