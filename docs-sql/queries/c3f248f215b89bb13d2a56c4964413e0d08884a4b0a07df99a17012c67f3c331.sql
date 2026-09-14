SET s3_base = 'https://datasets-documentation.s3.eu-west-3.amazonaws.com/';
CREATE TABLE aapl_stock (Date Date, Open Float32, High Float32, Low Float32, Close Float32, Volume Float32, OpenInt Int32)
    ENGINE = S3('aapl_stock.csv', NOSIGN, 'CSVWithNames');
