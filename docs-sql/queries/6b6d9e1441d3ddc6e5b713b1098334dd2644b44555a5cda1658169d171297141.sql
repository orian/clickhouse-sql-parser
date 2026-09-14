-- データベースを作成する
CREATE DATABASE db1;
CREATE DATABASE db2;

-- db1 にソーステーブルを作成する
CREATE TABLE db1.events (
    timestamp DateTime,
    event_type String,
    user_id UInt32
) ENGINE = MergeTree
ORDER BY timestamp;

-- db1.events を指す別名を db2 に作成する
CREATE TABLE db2.events_alias ENGINE = Alias('db1', 'events');

-- または database.table 形式を使用する
CREATE TABLE db2.events_alias2 ENGINE = Alias('db1.events');

-- どちらの別名も同じように機能する
INSERT INTO db2.events_alias VALUES (now(), 'click', 100);
SELECT * FROM db2.events_alias2;
