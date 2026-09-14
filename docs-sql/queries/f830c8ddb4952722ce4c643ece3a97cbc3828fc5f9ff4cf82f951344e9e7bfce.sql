-- 중복 제거된 posts 테이블 생성 
CREATE TABLE deduplicated_posts AS posts;

-- Materialized view 생성 및 매시간 실행 예약
CREATE MATERIALIZED VIEW deduplicated_posts_mv REFRESH EVERY 1 HOUR TO deduplicated_posts AS 
SELECT * FROM posts FINAL WHERE _peerdb_is_deleted=0 
