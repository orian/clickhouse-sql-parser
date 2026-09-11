     SELECT
         pid,
         state,
         age(now(), xact_start) AS transaction_duration,
         query AS current_query
     FROM
         pg_stat_activity
     WHERE
         xact_start IS NOT NULL
     ORDER BY
         age(now(), xact_start) DESC;
