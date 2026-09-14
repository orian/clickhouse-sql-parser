WITH
    (
        SELECT sum(cityHash64(*))
        FROM
        (
            SELECT *
            FROM numbers(10)
            ORDER BY number DESC
        )
    ) AS q1_resultset_hash,
    (
        SELECT sum(cityHash64(*))
        FROM
        (
            SELECT *
            FROM numbers(10)
            ORDER BY number ASC
        )
    ) AS q2_resultset_hash
SELECT q1_resultset_hash = q2_resultset_hash AS Q1_equals_Q2
FORMAT Pretty
