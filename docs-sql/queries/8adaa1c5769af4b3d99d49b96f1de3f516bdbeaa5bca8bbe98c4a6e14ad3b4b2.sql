SELECT
    m.name,
    m.id,
    g.movie_id,
    g.genre
FROM movies AS m
CROSS JOIN genres AS g
LIMIT 10;
