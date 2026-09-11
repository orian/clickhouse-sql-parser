CREATE TABLE events (
    id UInt32,
    data String
) ENGINE = MergeTree
ORDER BY id;

CREATE TABLE events_alias ENGINE = Alias('events');

-- Multiple inserts create multiple parts
INSERT INTO events_alias VALUES (1, 'data1');
INSERT INTO events_alias VALUES (2, 'data2');
INSERT INTO events_alias VALUES (3, 'data3');

-- Check parts count
SELECT count() FROM system.parts
WHERE database = currentDatabase()
AND table = 'events'
AND active;

-- Optimize through alias
OPTIMIZE TABLE events_alias FINAL;

-- Parts are merged in target table
SELECT count() FROM system.parts
WHERE database = currentDatabase()
AND table = 'events'
AND active;  -- Returns 1
