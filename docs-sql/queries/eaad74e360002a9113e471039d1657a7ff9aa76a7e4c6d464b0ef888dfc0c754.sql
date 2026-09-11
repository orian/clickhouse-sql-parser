SELECT count() FROM posts WHERE has(keywords, 'clickhouse'); -- 低速なフルテーブルスキャン - すべての投稿のすべてのキーワードをチェックする
