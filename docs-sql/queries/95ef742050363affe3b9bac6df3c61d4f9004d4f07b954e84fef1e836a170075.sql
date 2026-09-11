CREATE TABLE times (
    id     INT NOT NULL,
    months INT NOT NULL,
    days   INT NOT NULL
) WITH (copy_from = 's3://datasets-documentation/my-test-bucket-768/some_prefix/some_file_1.csv');
