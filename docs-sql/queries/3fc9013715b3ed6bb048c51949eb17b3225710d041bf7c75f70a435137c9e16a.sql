SET enable_materialized_cte = 1;

WITH
    b AS MATERIALIZED (SELECT uid FROM a),
    a AS MATERIALIZED (SELECT uid FROM users)
SELECT count() FROM b AS l LEFT SEMI JOIN b AS r ON l.uid = r.uid;
