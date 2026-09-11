SELECT d, extract(d, '([0-3])') AS res, toTypeName(res) FROM test;
