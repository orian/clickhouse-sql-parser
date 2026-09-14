CREATE TABLE events (
    date Date DEFAULT today(), 
    name String
) 
ENGINE = MergeTree
ORDER BY date;
