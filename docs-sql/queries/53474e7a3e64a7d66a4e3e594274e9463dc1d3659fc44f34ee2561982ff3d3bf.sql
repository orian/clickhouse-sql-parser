CREATE MATERIALIZED VIEW mv_vote_processor TO vote_aggregates
AS
SELECT
  post_id,
  -- sum 상태의 초기값 (upvote이면 1, 그 외에는 0)
  toUInt64(vote_type = 'upvote') AS upvotes,
  -- sum 상태의 초기값 (downvote이면 1, 그 외에는 0)
  toUInt64(vote_type = 'downvote') AS downvotes,
  -- sum 상태의 초기값 (upvote이면 1, downvote이면 -1)
  toInt64(vote_type) AS score
FROM raw_votes;
