SELECT 
    title,
    text,
    COUNT(DISTINCT concept) AS num_concepts_matched,
    MIN(distance) AS min_distance,
    AVG(distance) AS avg_distance
FROM (
    (
        SELECT title, text, 'Moon' AS concept,
               L2DistanceTransposed(qbit, (SELECT vector FROM dbpedia WHERE title = 'Moon'), 5) AS distance
        FROM dbpedia
        WHERE title != 'Moon'
        ORDER BY distance ASC
        LIMIT 1000
    )
    UNION ALL
    (
        SELECT title, text, 'Apollo 11' AS concept,
               L2DistanceTransposed(qbit, (SELECT vector FROM dbpedia WHERE title = 'Apollo 11'), 5) AS distance
        FROM dbpedia
        WHERE title != 'Apollo 11'
        ORDER BY distance ASC
        LIMIT 1000
    )
    UNION ALL
    (
        SELECT title, text, 'Space Shuttle' AS concept,
               L2DistanceTransposed(qbit, (SELECT vector FROM dbpedia WHERE title = 'Space Shuttle'), 5) AS distance
        FROM dbpedia
        WHERE title != 'Space Shuttle'
        ORDER BY distance ASC
        LIMIT 1000
    )
    UNION ALL
    (
        SELECT title, text, 'Astronaut' AS concept,
               L2DistanceTransposed(qbit, (SELECT vector FROM dbpedia WHERE title = 'Astronaut'), 5) AS distance
        FROM dbpedia
        WHERE title != 'Astronaut'
        ORDER BY distance ASC
        LIMIT 1000
    )
    UNION ALL
    (
        SELECT title, text, 'Rocket' AS concept,
               L2DistanceTransposed(qbit, (SELECT vector FROM dbpedia WHERE title = 'Rocket'), 5) AS distance
        FROM dbpedia
        WHERE title != 'Rocket'
        ORDER BY distance ASC
        LIMIT 1000
    )
)
WHERE title NOT IN ('Moon', 'Apollo 11', 'Space Shuttle', 'Astronaut', 'Rocket')
GROUP BY title, text
HAVING num_concepts_matched >= 3
ORDER BY num_concepts_matched DESC, min_distance ASC
LIMIT 10;
