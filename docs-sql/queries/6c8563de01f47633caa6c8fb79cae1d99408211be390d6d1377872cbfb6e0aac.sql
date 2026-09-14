INSERT INTO FUNCTION bigquery('my-project', 'my_dataset', 'my_table', '<access token>')
SELECT number AS id, toString(number) AS name FROM numbers(10);
