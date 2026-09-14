WITH [...] AS reference_vec
SELECT id
FROM vectors
ORDER BY cosineDistance(vec, reference_vec) ASC
LIMIT 10
SETTINGS vector_search_use_quantized_codes = 1;
