SELECT
    day_of_week,
    count() AS c
FROM git.commits
GROUP BY dayOfWeek(time) AS day_of_week
