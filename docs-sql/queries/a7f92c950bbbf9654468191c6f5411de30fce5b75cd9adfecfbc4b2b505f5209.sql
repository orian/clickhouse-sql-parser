SELECT count(*)
FROM s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/my-test-bucket-768/big_prefix/file-{1..4}.csv', NOSIGN, 'CSV', 'column1 UInt32, column2 UInt32, column3 UInt32');
