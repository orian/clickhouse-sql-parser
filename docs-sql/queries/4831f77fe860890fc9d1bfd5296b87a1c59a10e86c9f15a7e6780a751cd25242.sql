CREATE TABLE my_test_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(
    'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
    1,  -- データ生成のシード
    10  -- 異なるランダム値の数
)
LIMIT 100;  -- 100行

-- ステップ2: 新しいテーブルにクエリを実行する
SELECT * FROM my_test_table LIMIT 5;
