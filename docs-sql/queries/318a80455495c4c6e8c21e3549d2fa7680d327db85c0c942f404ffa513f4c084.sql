CREATE MATERIALIZED VIEW posts_final
REFRESH EVERY 10 second ENGINE = ReplacingMergeTree(_peerdb_version)
ORDER BY (owneruserid,id) -- chave de ordenação diferente, mas com sufixo da chave primária do Postgres
AS
SELECT * FROM posts FINAL 
WHERE _peerdb_is_deleted = 0; -- isso realiza a desduplicação
