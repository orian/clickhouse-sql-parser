    SELECT id, url 
    FROM laion_5b_10m
    ORDER BY cosineDistance( vector, (SELECT vector FROM laion_5b_10m WHERE id = 9999) ) ASC
    LIMIT 20

