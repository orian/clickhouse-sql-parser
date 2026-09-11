CREATE TABLE prod
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

CREATE TABLE staging
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

-- 初始数据
INSERT INTO prod VALUES (123, 'John', 33);
INSERT INTO prod VALUES (456, 'Ksenia', 48);
-- 加载数据
INSERT INTO staging VALUES (8811, 'Alice', 50);
INSERT INTO staging VALUES (8812, 'Bob', 23);

-- 验证导入
SELECT 'Staging count:', COUNT() FROM staging;
-- 移动分区
ALTER TABLE staging MOVE PARTITION tuple() TO TABLE prod; -- 原子操作

-- 检查数据
SELECT 'Prod count:', COUNT() FROM prod;
SELECT * FROM prod;
