INSERT INTO posts_with_duplicate_count SELECT
    posts.*,
    DuplicatePosts
FROM posts AS posts
LEFT JOIN
(
    SELECT PostId, countIf(LinkTypeId = 'Duplicate') AS DuplicatePosts
    FROM postlinks
    GROUP BY PostId
) AS postlinks ON posts.Id = postlinks.PostId
