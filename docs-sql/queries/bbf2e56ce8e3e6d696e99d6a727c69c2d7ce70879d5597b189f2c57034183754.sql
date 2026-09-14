-- (owneruserid, id)에 UNIQUE INDEX 생성
CREATE UNIQUE INDEX posts_unique_owneruserid_idx ON posts(owneruserid, id);
-- 이 인덱스를 사용하도록 REPLICA IDENTITY 설정
ALTER TABLE posts REPLICA IDENTITY USING INDEX posts_unique_owneruserid_idx;
