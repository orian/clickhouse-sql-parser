CREATE MATERIALIZED VIEW mv_vote_processor TO vote_aggregates
AS
SELECT
  post_id,
  -- sum 状态的初始值（upvote 时为 1，否则为 0）
  toUInt64(vote_type = 'upvote') AS upvotes,
  -- sum 状态的初始值（downvote 时为 1，否则为 0）
  toUInt64(vote_type = 'downvote') AS downvotes,
  -- sum 状态的初始值（upvote 时为 1，downvote 时为 -1）
  toInt64(vote_type) AS score
FROM raw_votes;
