SELECT word, sum(word_count) AS c
FROM bigquery('bigquery-public-data', 'samples', 'shakespeare', '<access token>')
GROUP BY word
ORDER BY c DESC
LIMIT 5;
