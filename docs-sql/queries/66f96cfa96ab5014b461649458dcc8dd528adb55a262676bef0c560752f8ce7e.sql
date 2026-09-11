  CREATE TABLE a (x String) ENGINE = Memory;
  INSERT INTO a VALUES ('ClickHouse');

  SELECT * FROM a AS a1 JOIN a AS a2 ON a1.x <=> a2.x;

  ┌─x──────────┬─a2.x───────┐
  │ ClickHouse │ ClickHouse │
  └────────────┴────────────┘
