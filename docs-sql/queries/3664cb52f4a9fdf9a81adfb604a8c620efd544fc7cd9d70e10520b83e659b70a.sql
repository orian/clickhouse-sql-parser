-- 创建数据库
CREATE DATABASE db1;
CREATE DATABASE db2;

-- 在 db1 中创建源表
CREATE TABLE db1.events (
    timestamp DateTime,
    event_type String,
    user_id UInt32
) ENGINE = MergeTree
ORDER BY timestamp;

-- 在 db2 中创建指向 db1.events 的别名
CREATE TABLE db2.events_alias ENGINE = Alias('db1', 'events');

-- 或使用 database.table 格式
CREATE TABLE db2.events_alias2 ENGINE = Alias('db1.events');

-- 两个别名的效果完全相同
INSERT INTO db2.events_alias VALUES (now(), 'click', 100);
SELECT * FROM db2.events_alias2;
