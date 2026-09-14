
        SELECT id, 
            AVG(value) OVER (ORDER BY id ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS frame1, 
            MAX(value) OVER (ORDER BY id ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS frame2,
            sipHash64(frame1),
            sipHash64(frame2)
        FROM window_test
    