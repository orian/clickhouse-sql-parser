-- 浏览量最高的帖子
SELECT
    sum(viewcount) AS viewcount,
    owneruserid
FROM posts_view
WHERE owneruserid > 0
GROUP BY owneruserid
ORDER BY viewcount DESC
LIMIT 10
