-- Создать UNIQUE INDEX для (owneruserid, id)
CREATE UNIQUE INDEX posts_unique_owneruserid_idx ON posts(owneruserid, id);
-- Задать REPLICA IDENTITY для использования этого индекса
ALTER TABLE posts REPLICA IDENTITY USING INDEX posts_unique_owneruserid_idx;
