CREATE TABLE posts_with_duplicate_count
(
  `Id` Int32 CODEC(Delta(4), ZSTD(1)),
   ... -다른 컬럼들
   `DuplicatePosts` UInt16
) ENGINE = MergeTree
ORDER BY (PostTypeId, toDate(CreationDate), CommentCount)
