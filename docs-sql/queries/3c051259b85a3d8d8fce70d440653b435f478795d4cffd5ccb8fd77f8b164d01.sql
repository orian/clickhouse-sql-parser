-- 创建去重后的 posts 表 
CREATE TABLE deduplicated_posts AS posts;

-- 创建 Materialized view 并设置为每小时刷新一次
CREATE MATERIALIZED VIEW deduplicated_posts_mv REFRESH EVERY 1 HOUR TO deduplicated_posts AS 
SELECT * FROM posts FINAL WHERE _peerdb_is_deleted=0 
