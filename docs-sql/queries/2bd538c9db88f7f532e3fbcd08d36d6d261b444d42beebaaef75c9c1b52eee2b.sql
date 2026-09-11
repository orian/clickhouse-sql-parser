COPY users TO 'file:///tmp/users.parquet' (
    structure 'id Int64, name String, age Nullable(UInt8), attributes JSON'
);
