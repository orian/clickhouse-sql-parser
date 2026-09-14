CREATE MATERIALIZED VIEW posts_final
REFRESH EVERY 10 second ENGINE = ReplacingMergeTree(_peerdb_version)
ORDER BY (owneruserid,id) -- 異なる順序キー（接尾辞付きのPostgres主キーを含む）
AS
SELECT * FROM posts FINAL 
WHERE _peerdb_is_deleted = 0; -- これによりdeduplicationを実行する
