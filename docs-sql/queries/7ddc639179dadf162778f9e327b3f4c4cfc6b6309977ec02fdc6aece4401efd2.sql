INSERT INTO test_db.test_table SELECT *
FROM generateRandom('key Int, value String, array Array(String)')
LIMIT 1000
