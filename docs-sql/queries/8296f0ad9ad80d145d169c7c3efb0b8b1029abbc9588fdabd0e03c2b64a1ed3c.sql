    CREATE MATERIALIZED VIEW github_out_mv TO github_out_queue AS
    SELECT file_time, event_type, actor_login, repo_name,
           created_at, updated_at, action, comment_id, path,
           ref, ref_type, creator_user_login, number, title,
           labels, state, assignee, assignees, closed_at, merged_at,
           merge_commit_sha, requested_reviewers, merged_by,
           review_comments, member_login
    FROM github
    FORMAT JsonEachRow;
