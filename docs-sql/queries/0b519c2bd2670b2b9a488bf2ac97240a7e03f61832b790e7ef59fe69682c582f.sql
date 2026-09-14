CREATE TABLE big_table (name String, value UInt32)
    ENGINE = S3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/aapl_stock.csv', NOSIGN, 'CSVWithNames');
