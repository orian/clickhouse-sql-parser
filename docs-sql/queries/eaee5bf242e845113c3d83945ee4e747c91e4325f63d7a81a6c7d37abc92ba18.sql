CREATE MATERIALIZED VIEW github_mv TO github AS
SELECT *
FROM github_queue;
