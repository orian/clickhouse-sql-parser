SELECT * 
FROM file(data.parquet, ParquetMetadata) 
FORMAT PrettyJSONEachRow
