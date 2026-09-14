SELECT
  post_id,
  sum(upvotes) AS total_upvotes,
  sum(downvotes) AS total_downvotes,
  sum(score) AS total_score
FROM vote_aggregates -- 대상 테이블 쿼리
GROUP BY post_id
ORDER BY post_id ASC;
