    SELECT id, title, text
    FROM hackernews
    ORDER BY cosineDistance( vector, <search vector>)
    LIMIT 10

