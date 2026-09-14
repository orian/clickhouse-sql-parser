CREATE MATERIALIZED VIEW posts_final
REFRESH EVERY 10 second ENGINE = ReplacingMergeTree(_peerdb_version)
ORDER BY (owneruserid,id) -- 不同的排序键，但带有后缀的 postgres 主键
AS
SELECT * FROM posts FINAL 
WHERE _peerdb_is_deleted = 0; -- 此处执行去重
