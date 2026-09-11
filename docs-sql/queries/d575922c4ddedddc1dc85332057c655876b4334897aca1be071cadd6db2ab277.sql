CREATE TABLE comments
(
    `Id` UInt32,
    `PostId` UInt32,
    `Score` UInt16,
    `Text` String,
    `CreationDate` DateTime64(3, 'UTC'),
    `UserId` Int32,
    `UserDisplayName` LowCardinality(String),
    PROJECTION comments_user_id
    (
    SELECT *
    ORDER BY UserId
    )
)
ENGINE = MergeTree
ORDER BY PostId
