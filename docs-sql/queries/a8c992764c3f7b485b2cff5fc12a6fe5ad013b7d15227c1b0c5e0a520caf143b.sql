-- Cria um UNIQUE INDEX em (owneruserid, id)
CREATE UNIQUE INDEX posts_unique_owneruserid_idx ON posts(owneruserid, id);
-- Define a REPLICA IDENTITY para usar este índice
ALTER TABLE posts REPLICA IDENTITY USING INDEX posts_unique_owneruserid_idx;
