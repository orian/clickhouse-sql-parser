CREATE TABLE my_test_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(
    'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
    1,  -- seed for data generation
    10  -- number of different random values
)
LIMIT 100;  -- 100 rows

-- Step 2: Query your new table
SELECT * FROM my_test_table LIMIT 5;
