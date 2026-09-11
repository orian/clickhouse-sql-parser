INSERT INTO hackernews SELECT *
FROM url('https://datasets-documentation.s3.eu-west-3.amazonaws.com/hackernews/hacknernews.csv.gz', 'CSVWithNames')
