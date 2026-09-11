-- The output will be delayed, with a 0.5-second pause between each row.
SELECT number, sleepEachRow(0.5) FROM system.numbers LIMIT 5;
