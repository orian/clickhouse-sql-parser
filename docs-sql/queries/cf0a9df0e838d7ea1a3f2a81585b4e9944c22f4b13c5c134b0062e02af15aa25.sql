WITH
    (
        SELECT sum(cityHash64(*))
        FROM
        (
            -- your query 1 here
            -- SELECT ...
        )
    ) AS q1_resultset_hash,
    (
        SELECT sum(cityHash64(*))
        FROM
        (
            -- your query 2 here
            -- SELECT ...
        )
    ) AS q2_resultset_hash
SELECT equals(q1_resultset_hash,q2_resultset_hash) as Q1_equals_Q2
