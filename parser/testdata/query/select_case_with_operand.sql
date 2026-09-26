SELECT CASE number WHEN 1 THEN number + 2 ELSE number * 2 END FROM numbers(3);
SELECT CASE number % 3 WHEN 0 THEN 'zero' WHEN 1 THEN 'one' ELSE 'two' END AS name FROM numbers(6);
