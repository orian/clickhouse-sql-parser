SELECT id, url 
FROM laion_5b_10m
ORDER BY cosineDistance( vector, (SELECT vector FROM laion_5b_10m WHERE id = 9999) ) ASC
LIMIT 20

The vector in the row with id = 9999 is the embedding for an image of a Deli restaurant.
