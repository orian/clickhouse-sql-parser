-- 首先创建表并插入一些数据
CREATE TABLE hilbert_numbers(
    n1 UInt32,
    n2 UInt32
)
ENGINE=MergeTree()
ORDER BY n1;
insert into hilbert_numbers (*) values(1, 2);

-- 使用列名而非常量作为函数参数
SELECT hilbertEncode(n1, n2) FROM hilbert_numbers;
