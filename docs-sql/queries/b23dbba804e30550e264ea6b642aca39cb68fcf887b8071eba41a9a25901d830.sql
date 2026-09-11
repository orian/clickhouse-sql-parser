-- (owneruserid, id) に UNIQUE INDEX を作成する
CREATE UNIQUE INDEX posts_unique_owneruserid_idx ON posts(owneruserid, id);
-- このインデックスを使用するように REPLICA IDENTITY を設定する
ALTER TABLE posts REPLICA IDENTITY USING INDEX posts_unique_owneruserid_idx;
