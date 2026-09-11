-- Create databases
CREATE DATABASE db1;
CREATE DATABASE db2;

-- Create source table in db1
CREATE TABLE db1.events (
    timestamp DateTime,
    event_type String,
    user_id UInt32
) ENGINE = MergeTree
ORDER BY timestamp;

-- Create alias in db2 pointing to db1.events
CREATE TABLE db2.events_alias ENGINE = Alias('db1', 'events');

-- Or using database.table format
CREATE TABLE db2.events_alias2 ENGINE = Alias('db1.events');

-- Both aliases work identically
INSERT INTO db2.events_alias VALUES (now(), 'click', 100);
SELECT * FROM db2.events_alias2;
