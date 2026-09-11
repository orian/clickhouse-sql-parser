SELECT
 UserId,
 argMax(description, event_time) AS last_description,
 argMax(activity_type, event_time) AS activity_type,
    max(event_time) AS last_activity
FROM
(
    SELECT
 UserId,
 CreationDate AS event_time,
        Text AS description,
        'comment' AS activity_type
    FROM stackoverflow.comments
    UNION ALL
    SELECT
 UserId,
        Date AS event_time,
        Name AS description,
        'badge' AS activity_type
    FROM stackoverflow.badges
)
GROUP BY UserId
ORDER BY last_activity DESC
LIMIT 10
