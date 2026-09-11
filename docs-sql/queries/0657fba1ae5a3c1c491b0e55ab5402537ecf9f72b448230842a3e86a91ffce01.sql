CREATE TABLE reviews () WITH (
    copy_from = 's3://datasets-documentation/amazon_reviews/amazon_reviews_2015.snappy.parquet'
);
