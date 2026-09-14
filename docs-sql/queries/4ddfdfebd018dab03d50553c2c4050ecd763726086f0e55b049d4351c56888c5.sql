-- まずテーブルを作成してデータを挿入する
CREATE TABLE hilbert_numbers(
    n1 UInt32,
    n2 UInt32
)
ENGINE=MergeTree()
ORDER BY n1;
insert into hilbert_numbers (*) values(1, 2);

-- 関数の引数として定数の代わりにカラム名を使用する
SELECT hilbertEncode(n1, n2) FROM hilbert_numbers;
