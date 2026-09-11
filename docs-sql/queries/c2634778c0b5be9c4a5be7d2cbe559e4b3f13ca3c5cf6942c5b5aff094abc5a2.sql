SELECT * FROM fuzzJSON(json_nc, json_str='{"name" : "FuzzJSON"}', random_seed=1337, malform_output=true) LIMIT 3;
