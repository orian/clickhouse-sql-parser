    SELECT
        _file,
    	caption,
    	cosineDistance(
            -- Um embedding da sua imagem de cachorro de "entrada"
            [0.5736801028251648, 0.2516217529773712, ...,  -0.6825592517852783],
            image_embedding
        ) AS score
    FROM images
    ORDER BY score ASC
    LIMIT 10
