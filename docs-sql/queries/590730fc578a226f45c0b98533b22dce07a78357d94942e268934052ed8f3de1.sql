CREATE MATERIALIZED VIEW posts_final
REFRESH EVERY 10 second ENGINE = ReplacingMergeTree(_peerdb_version)
ORDER BY (owneruserid,id) -- clave de ordenación diferente pero con sufijo de clave primaria de Postgres
AS
SELECT * FROM posts FINAL 
WHERE _peerdb_is_deleted = 0; -- esto realiza la deduplicación
