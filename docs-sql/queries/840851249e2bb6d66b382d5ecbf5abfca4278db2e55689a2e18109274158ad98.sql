CREATE MATERIALIZED VIEW mv_vote_processor TO vote_aggregates
AS
SELECT
  post_id,
  -- Valor inicial para o estado de soma (1 se upvote, 0 caso contrário)
  toUInt64(vote_type = 'upvote') AS upvotes,
  -- Valor inicial para o estado de soma (1 se downvote, 0 caso contrário)
  toUInt64(vote_type = 'downvote') AS downvotes,
  -- Valor inicial para o estado de soma (1 para upvote, -1 para downvote)
  toInt64(vote_type) AS score
FROM raw_votes;
