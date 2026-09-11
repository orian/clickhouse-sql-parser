-- ソーステーブルを作成
CREATE TABLE source_data (
    id UInt32,
    name String,
    value Float64
) ENGINE = MergeTree
ORDER BY id;

-- データを挿入
INSERT INTO source_data VALUES (1, 'one', 10.1), (2, 'two', 20.2);

-- エイリアスを作成
CREATE TABLE data_alias ENGINE = Alias('source_data');

-- エイリアス経由でクエリを実行
SELECT * FROM data_alias;
