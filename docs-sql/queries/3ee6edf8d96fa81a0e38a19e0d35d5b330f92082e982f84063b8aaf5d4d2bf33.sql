WITH
    changes AS
    (
        SELECT
            path,
            author,
            commit_hash,
            max_time,
            type,
            num_added,
            num_deleted,
            sum(num_added - num_deleted) OVER (PARTITION BY path ORDER BY max_time ASC) AS current_size,
            if(current_size > 0, num_added / current_size, 0) AS percent_add,
            if(current_size > 0, num_deleted / current_size, 0) AS percent_delete
        FROM
        (
            SELECT
                path,
                any(author) AS author,
                max(time) AS max_time,
                commit_hash,
                any(file_lines_added) AS num_added,
                any(file_lines_deleted) AS num_deleted,
                any(file_change_type) AS type
            FROM git.line_changes
            WHERE (file_change_type IN ('Add', 'Modify')) AND (file_extension IN ('h', 'cpp', 'sql'))
            GROUP BY
                path,
                commit_hash
            ORDER BY
                path ASC,
                max_time ASC
        )
    ),
    rewrites AS
    (
        SELECT
            *,
            any(max_time) OVER (PARTITION BY path ORDER BY max_time ASC ROWS BETWEEN 1 PRECEDING AND CURRENT ROW) AS previous_rewrite,
            dateDiff('day', previous_rewrite, max_time) AS rewrite_days,
            any(author) OVER (PARTITION BY path ORDER BY max_time ASC ROWS BETWEEN 1 PRECEDING AND CURRENT ROW) AS prev_author
        FROM changes
        WHERE (type = 'Modify') AND (percent_add >= 0.5) AND (percent_delete >= 0.5) AND (current_size > 50)
    )
SELECT
    prev_author,
    avg(rewrite_days) AS c,
    uniq(path) AS num_files
FROM rewrites
GROUP BY prev_author
HAVING num_files > 2
ORDER BY c DESC
LIMIT 10
