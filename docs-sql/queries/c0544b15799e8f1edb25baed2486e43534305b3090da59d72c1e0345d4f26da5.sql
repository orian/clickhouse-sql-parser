-- まずテーブルを作成し、データを挿入する
CREATE TABLE morton_numbers(
    n1 UInt32,
    n2 UInt32,
    n3 UInt16,
    n4 UInt16,
    n5 UInt8,
    n6 UInt8,
    n7 UInt8,
    n8 UInt8
)
ENGINE=MergeTree()
ORDER BY n1;
INSERT INTO morton_numbers (*) values(1, 2, 3, 4, 5, 6, 7, 8);

-- 関数の引数として定数の代わりにカラム名を使用する
SELECT mortonEncode(n1, n2, n3, n4, n5, n6, n7, n8) FROM morton_numbers;
