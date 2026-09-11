CREATE MATERIALIZED VIEW posts_final
REFRESH EVERY 10 second ENGINE = ReplacingMergeTree(_peerdb_version)
ORDER BY (owneruserid,id) -- 다른 정렬 키이지만 postgres 기본 키에 접미사가 추가된 형태
AS
SELECT * FROM posts FINAL 
WHERE _peerdb_is_deleted = 0; -- 중복 제거 수행
