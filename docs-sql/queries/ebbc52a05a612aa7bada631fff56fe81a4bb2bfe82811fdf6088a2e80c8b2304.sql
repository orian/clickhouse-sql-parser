CREATE MATERIALIZED VIEW posts_final
REFRESH EVERY 10 second ENGINE = ReplacingMergeTree(_peerdb_version)
ORDER BY (owneruserid,id) -- другой ключ сортировки, но с суффиксом первичного ключа postgres
AS
SELECT * FROM posts FINAL 
WHERE _peerdb_is_deleted = 0; -- выполняет дедупликацию
