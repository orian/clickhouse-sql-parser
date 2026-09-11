-- `hits` テーブルから `Title` カラムに `hello` というテキストを含むすべての行を削除する
DELETE FROM hits WHERE Title LIKE '%hello%';
