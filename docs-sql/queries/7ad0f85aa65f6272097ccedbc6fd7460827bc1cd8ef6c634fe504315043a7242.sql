CREATE TABLE times (
    id     INT PRIMARY KEY,
    months INT NOT NULL,
    days   INT NOT NULL
);

COPY times FROM 's3://datasets-documentation/my-test-bucket-768/some_prefix/some_file_1.csv';
