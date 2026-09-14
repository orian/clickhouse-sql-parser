SET enable_materialized_cte = 1;

CREATE TABLE users (uid Int16, name String, age Int16) ENGINE = Memory;
INSERT INTO users VALUES (1231, 'John', 33), (6666, 'Ksenia', 48), (8888, 'Alice', 50);

WITH
    a AS MATERIALIZED (SELECT * FROM users WHERE name = 'Alice')
SELECT count() FROM a AS l JOIN a AS r ON l.uid = r.uid;
