SELECT
    url,
    caption
FROM laion
ORDER BY cosineDistance(image_embedding, encode_image('/path/to/your/image')) ASC
LIMIT 10
