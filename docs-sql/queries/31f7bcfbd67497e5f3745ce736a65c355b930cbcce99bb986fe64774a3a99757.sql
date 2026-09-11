-- 重複排除済みのpostsテーブルを作成する 
CREATE TABLE deduplicated_posts AS posts;

-- materialized viewを作成し、1時間ごとに実行するようスケジュールする
CREATE MATERIALIZED VIEW deduplicated_posts_mv REFRESH EVERY 1 HOUR TO deduplicated_posts AS 
SELECT * FROM posts FINAL WHERE _peerdb_is_deleted=0 
