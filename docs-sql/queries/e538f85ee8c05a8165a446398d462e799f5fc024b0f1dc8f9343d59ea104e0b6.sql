    DROP TABLE IF EXISTS hackernews;

    CREATE TABLE hackernews
    (
        `id` UInt64,
        `deleted` UInt8,
        `type` String,
        `author` String,
        `timestamp` DateTime,
        `comment` String,
        `dead` UInt8,
        `parent` UInt64,
        `poll` UInt64,
        `children` Array(UInt32),
        `url` String,
        `score` UInt32,
        `title` String,
        `parts` Array(UInt32),
        `descendants` UInt32
    )
    ENGINE = MergeTree
    ORDER BY (type, author);

    INSERT INTO hackernews
    SELECT * FROM s3(
            'https://datasets-documentation.s3.eu-west-3.amazonaws.com/hackernews/hacknernews.parquet',
            NOSIGN,
            'Parquet',
            'id UInt64,
             deleted UInt8,
             type String,
             by String,
             time DateTime,
             text String,
             dead UInt8,
             parent UInt64,
             poll UInt64,
             kids Array(UInt32),
             url String,
             score UInt32,
             title String,
             parts Array(UInt32),
             descendants UInt32');
