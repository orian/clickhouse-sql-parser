
        SELECT
            l.number,
            r.s0,
            r.s1,
            r.s2,
            r.s3,
            r.u0,
            r.u1,
            r.u2,
            r.u3
        FROM numbers(50000000) AS l
        ANY LEFT JOIN tab AS r ON 1
        FORMAT Null
    