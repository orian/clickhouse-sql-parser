-- Origin SQL:
SELECT c0 FROM remote('localhost', currentDatabase(), 't0') AS tx INNER JOIN t1 USING (c0);
SELECT * FROM merge(currentDatabase(), '^t');
SELECT * FROM file(currentDatabase() || '_1.arrow', 'Arrow');
SELECT * FROM numbers(1 + 2);
SELECT * FROM numbers(-1);
SELECT * FROM numbers(10, 5 * 2);
SELECT * FROM numbers(toUInt64(3));
SELECT * FROM s3(concat('http://x/', 'y.csv'), 'CSV');
SELECT * FROM s3('u', format = lower('CSV'));


-- Format SQL:
SELECT c0 FROM remote('localhost', currentDatabase(), 't0') AS tx INNER JOIN t1 USING (c0);
SELECT * FROM merge(currentDatabase(), '^t');
SELECT * FROM file(currentDatabase() || '_1.arrow', 'Arrow');
SELECT * FROM numbers(1 + 2);
SELECT * FROM numbers(-1);
SELECT * FROM numbers(10, 5 * 2);
SELECT * FROM numbers(toUInt64(3));
SELECT * FROM s3(concat('http://x/', 'y.csv'), 'CSV');
SELECT * FROM s3('u', format=lower('CSV'));
