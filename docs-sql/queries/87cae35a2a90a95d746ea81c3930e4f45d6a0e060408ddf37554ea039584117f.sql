SELECT
  post_id,
  sum(upvotes) AS total_upvotes,
  sum(downvotes) AS total_downvotes,
  sum(score) AS total_score
FROM vote_aggregates -- Запрос к целевой таблице
GROUP BY post_id
ORDER BY post_id ASC;
