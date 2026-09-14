SET enable_materialized_cte = 1;

WITH
    a AS MATERIALIZED (SELECT uid, name FROM users),
    b AS MATERIALIZED (SELECT uid FROM a)
SELECT count() FROM b AS l LEFT SEMI JOIN b AS r ON l.uid = r.uid;
