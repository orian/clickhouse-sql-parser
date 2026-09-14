-- 创建源表
CREATE TABLE source_data (
    id UInt32,
    name String,
    value Float64
) ENGINE = MergeTree
ORDER BY id;

-- 插入一些数据
INSERT INTO source_data VALUES (1, 'one', 10.1), (2, 'two', 20.2);

-- 创建别名
CREATE TABLE data_alias ENGINE = Alias('source_data');

-- 通过别名查询
SELECT * FROM data_alias;
