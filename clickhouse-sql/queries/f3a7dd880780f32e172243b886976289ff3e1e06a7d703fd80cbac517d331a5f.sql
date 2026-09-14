
        SELECT id AS key, 
            sipHash64(sum(frame)) AS value 
        FROM (
            SELECT id, 
                AVG(value) OVER (ORDER BY id ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS frame 
            FROM window_test) 
        GROUP BY key 
        ORDER BY key, value
    