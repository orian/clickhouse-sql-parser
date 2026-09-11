SELECT toYear(CreationDate) AS year, sum(AnswerCount) AS total_answers
FROM posts_no_part
FINAL
GROUP BY year
ORDER BY year ASC
