CREATE MATERIALIZED VIEW mv_vote_processor TO vote_aggregates
AS
SELECT
  post_id,
  -- sum状態の初期値（upvoteの場合は1、それ以外は0）
  toUInt64(vote_type = 'upvote') AS upvotes,
  -- sum状態の初期値（downvoteの場合は1、それ以外は0）
  toUInt64(vote_type = 'downvote') AS downvotes,
  -- sum状態の初期値（upvoteの場合は1、downvoteの場合は-1）
  toInt64(vote_type) AS score
FROM raw_votes;
