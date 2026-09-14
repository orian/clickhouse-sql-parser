-- まず、データの保存に使用されているデータベースを特定します。
SELECT * FROM [database].connect_state

-- トピックとパーティションに一致するキーを特定します。
ALTER TABLE [database].connect_state DELETE WHERE key = [keyname]
