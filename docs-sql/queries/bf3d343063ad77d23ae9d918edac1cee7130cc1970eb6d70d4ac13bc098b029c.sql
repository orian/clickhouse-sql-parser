CREATE TABLE my_test_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(
    'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
    1,  -- 数据生成的种子值
    10  -- 不同随机值的数量
)
LIMIT 100;  -- 100 行

-- 第 2 步：查询新表
SELECT * FROM my_test_table LIMIT 5;
