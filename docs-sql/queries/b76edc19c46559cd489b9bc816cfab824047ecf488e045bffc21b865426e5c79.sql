-- 데이터베이스 생성
CREATE DATABASE db1;
CREATE DATABASE db2;

-- db1에 원본 테이블 생성
CREATE TABLE db1.events (
    timestamp DateTime,
    event_type String,
    user_id UInt32
) ENGINE = MergeTree
ORDER BY timestamp;

-- db1.events를 가리키는 별칭을 db2에 생성
CREATE TABLE db2.events_alias ENGINE = Alias('db1', 'events');

-- 또는 database.table 포맷 사용
CREATE TABLE db2.events_alias2 ENGINE = Alias('db1.events');

-- 두 별칭 모두 동일하게 작동
INSERT INTO db2.events_alias VALUES (now(), 'click', 100);
SELECT * FROM db2.events_alias2;
