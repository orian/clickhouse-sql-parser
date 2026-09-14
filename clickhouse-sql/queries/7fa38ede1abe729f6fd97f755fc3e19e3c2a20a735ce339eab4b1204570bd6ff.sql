
        SELECT id, 
            AVG(value) OVER (PARTITION by partition ORDER BY id ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS frame 
        FROM window_test 
        ORDER BY id
    