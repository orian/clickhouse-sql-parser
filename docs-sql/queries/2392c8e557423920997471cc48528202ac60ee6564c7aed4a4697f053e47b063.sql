CREATE MATERIALIZED VIEW mv_vote_processor TO vote_aggregates
AS
SELECT
  post_id,
  -- Начальное значение для состояния суммы (1 если upvote, иначе 0)
  toUInt64(vote_type = 'upvote') AS upvotes,
  -- Начальное значение для состояния суммы (1 если downvote, иначе 0)
  toUInt64(vote_type = 'downvote') AS downvotes,
  -- Начальное значение для состояния суммы (1 для upvote, -1 для downvote)
  toInt64(vote_type) AS score
FROM raw_votes;
