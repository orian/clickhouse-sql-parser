SELECT * FROM fuzzQuery('SELECT materialize(\'a\' AS key) GROUP BY key') LIMIT 2;
