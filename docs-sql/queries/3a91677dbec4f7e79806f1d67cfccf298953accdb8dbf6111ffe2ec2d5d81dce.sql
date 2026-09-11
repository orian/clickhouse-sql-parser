SELECT substring(a, 4, 2), substring(substring(a, 1, 2), 1, count(b)) FROM t LIMIT 2 BY ALL;
