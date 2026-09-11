
        SELECT * FROM
            (SELECT EventTime,
                count(*) OVER (ORDER BY EventTime ASC
                    RANGE BETWEEN 10 PRECEDING AND 10 FOLLOWING) AS c
            FROM hits_10m_single)
        FORMAT Null
    