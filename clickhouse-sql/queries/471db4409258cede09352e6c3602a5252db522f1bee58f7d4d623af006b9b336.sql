
        SELECT count()
        FROM
        (
            SELECT
                t1.col_1 AS col_1,
                t1.col_2 AS col_2,
                t1.col_3 AS col_3,
                t2.col_5 AS col_4
            FROM
            (
                SELECT *
                FROM VALUES('col_1 DateTime64(6), col_2 LowCardinality(String), col_3 UUID, col_10 UUID',
                    ('1993-01-01', 'x', '75d076fd-84f3-4684-b830-090eca3d33a3', '75d076fd-84f3-4684-b830-090eca3d33a3'))
            ) AS t1
            ANY LEFT JOIN
            (
                SELECT *
                FROM VALUES('id UUID, col_5 String', ('75d076fd-84f3-4684-b830-090eca3d33a3', 'x'))
            ) AS t2 ON t2.id = t1.col_3
            ANY LEFT JOIN
            (
                SELECT *
                FROM VALUES('id UUID', '75d076fd-84f3-4684-b830-090eca3d33a3')
            ) AS t3 ON t3.id = t1.col_10
            WHERE col_2 IN ('qqq')
        ) AS t
        WHERE (col_1 > 0) AND (length(col_4) > 0)
    