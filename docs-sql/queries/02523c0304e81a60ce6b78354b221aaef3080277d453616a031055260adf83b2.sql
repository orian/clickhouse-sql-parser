CREATE MATERIALIZED VIEW mv_vote_processor TO vote_aggregates
AS
SELECT
  post_id,
  -- Initial value for sum state (1 if upvote, 0 otherwise)
  toUInt64(vote_type = 'upvote') AS upvotes,
  -- Initial value for sum state (1 if downvote, 0 otherwise)
  toUInt64(vote_type = 'downvote') AS downvotes,
  -- Initial value for sum state (1 for upvote, -1 for downvote)
  toInt64(vote_type) AS score
FROM raw_votes;
