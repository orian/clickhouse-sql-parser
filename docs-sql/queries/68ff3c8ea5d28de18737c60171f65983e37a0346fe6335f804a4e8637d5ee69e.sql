SELECT
    _file,
	caption,
	cosineDistance(
        -- An embedding of your "input" dog picture
        [0.5736801028251648, 0.2516217529773712, ...,  -0.6825592517852783],
        image_embedding
    ) AS score
FROM images
ORDER BY score ASC
LIMIT 10
