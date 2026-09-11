SELECT path, hits
FROM file('data.avro', Avro)
ORDER BY hits DESC
LIMIT 5;
