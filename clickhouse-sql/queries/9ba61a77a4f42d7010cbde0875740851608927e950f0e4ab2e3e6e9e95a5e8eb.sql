
        SELECT
            toDate(dt) AS dt,
            count()
        FROM test_pk_long
        WHERE (dt >= '2001-01-02') AND (dt <= '2001-01-11') AND (key = '1')
        GROUP BY dt
        ORDER BY dt ASC;
    