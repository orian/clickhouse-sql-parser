    CREATE TABLE users
    ENGINE = MergeTree ORDER BY username
    EMPTY AS
    SELECT * FROM s3('https://your-bucket.s3.amazonaws.com/unload/users/*', '<aws_access_key>', '<aws_secret_access_key>', 'CSV')
