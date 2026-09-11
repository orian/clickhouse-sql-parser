SET url_base = 'https://raw.githubusercontent.com/ClickHouse/ClickHouse/master/';
SELECT * FROM url('tests/queries/0_stateless/data_csv/data.csv', CSV) LIMIT 3;
