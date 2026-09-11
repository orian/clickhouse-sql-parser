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

-- 初期データ
INSERT INTO prod VALUES (123, 'John', 33);
INSERT INTO prod VALUES (456, 'Ksenia', 48);
-- データの読み込み
INSERT INTO staging VALUES (8811, 'Alice', 50);
INSERT INTO staging VALUES (8812, 'Bob', 23);

-- インポートの検証
SELECT 'Staging count:', COUNT() FROM staging;
-- パーティションの移動
ALTER TABLE staging MOVE PARTITION tuple() TO TABLE prod; -- アトミック操作

-- データの確認
SELECT 'Prod count:', COUNT() FROM prod;
SELECT * FROM prod;
