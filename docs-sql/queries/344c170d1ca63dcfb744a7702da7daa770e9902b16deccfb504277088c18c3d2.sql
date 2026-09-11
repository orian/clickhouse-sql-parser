CREATE MATERIALIZED VIEW mv_vote_processor TO vote_aggregates
AS
SELECT
  post_id,
  -- Valor inicial para el estado de suma (1 si es upvote, 0 en caso contrario)
  toUInt64(vote_type = 'upvote') AS upvotes,
  -- Valor inicial para el estado de suma (1 si es downvote, 0 en caso contrario)
  toUInt64(vote_type = 'downvote') AS downvotes,
  -- Valor inicial para el estado de suma (1 para upvote, -1 para downvote)
  toInt64(vote_type) AS score
FROM raw_votes;
