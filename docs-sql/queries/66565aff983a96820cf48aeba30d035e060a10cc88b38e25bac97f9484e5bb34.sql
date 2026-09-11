INSERT INTO amazon.amazon_reviews SELECT *
FROM s3Cluster('default', 
'https://datasets-documentation.s3.eu-west-3.amazonaws.com/amazon_reviews/amazon_reviews_*.snappy.parquet', NOSIGN)
