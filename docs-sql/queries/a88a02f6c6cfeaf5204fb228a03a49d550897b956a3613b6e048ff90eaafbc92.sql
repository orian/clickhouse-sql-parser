-- 在 (owneruserid, id) 上创建唯一索引
CREATE UNIQUE INDEX posts_unique_owneruserid_idx ON posts(owneruserid, id);
-- 将 REPLICA IDENTITY 设置为使用此索引
ALTER TABLE posts REPLICA IDENTITY USING INDEX posts_unique_owneruserid_idx;
