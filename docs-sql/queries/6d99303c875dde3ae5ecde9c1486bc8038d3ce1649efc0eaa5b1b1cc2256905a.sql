EXPLAIN distributed=1 SELECT * FROM remote('127.0.0.{1,2}', numbers(2)) WHERE number = 1;
