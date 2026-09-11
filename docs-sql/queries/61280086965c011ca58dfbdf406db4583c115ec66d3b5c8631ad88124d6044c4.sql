CREATE TABLE stackoverflow.comments
(
    `Id` UInt32,
    `PostId` UInt32,
    `Score` UInt16,
    `Text` String,
    `CreationDate` DateTime64(3, 'UTC'),
    `UserId` Int32,
    `UserDisplayName` LowCardinality(String)
)
ENGINE = MergeTree
ORDER BY CreationDate

INSERT INTO stackoverflow.comments SELECT * FROM s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/stackoverflow/parquet/comments/*.parquet', NOSIGN)
