SELECT groupFormat('JSONEachRow')(if(number = 0, NULL, number))
FROM numbers(3);
-- {"c1":1}
-- {"c1":2}
