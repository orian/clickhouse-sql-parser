-- Создать таблицу дедуплицированных постов 
CREATE TABLE deduplicated_posts AS posts;

-- Создать materialized view и запланировать запуск каждый час
CREATE MATERIALIZED VIEW deduplicated_posts_mv REFRESH EVERY 1 HOUR TO deduplicated_posts AS 
SELECT * FROM posts FINAL WHERE _peerdb_is_deleted=0 
