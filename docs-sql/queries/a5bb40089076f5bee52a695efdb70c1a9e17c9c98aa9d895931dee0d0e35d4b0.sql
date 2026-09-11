-- Crear tabla de posts deduplicados 
CREATE TABLE deduplicated_posts AS posts;

-- Crear la vista materializada y programar su ejecución cada hora
CREATE MATERIALIZED VIEW deduplicated_posts_mv REFRESH EVERY 1 HOUR TO deduplicated_posts AS 
SELECT * FROM posts FINAL WHERE _peerdb_is_deleted=0 
