INSERT INTO wikistat 
SELECT *
FROM s3('https://ClickHouse-public-datasets.s3.amazonaws.com/wikistat/partitioned/wikistat*.native.zst', NOSIGN)
LIMIT 1e9;
