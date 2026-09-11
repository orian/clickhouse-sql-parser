CREATE OR REPLACE TABLE mytable
(id UInt32, message String)
ENGINE MergeTree()
ORDER BY (id)
