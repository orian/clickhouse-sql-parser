CREATE NAMED COLLECTION json_fuzzer AS json_str='{}';
SELECT * FROM fuzzJSON(json_fuzzer) LIMIT 3;
