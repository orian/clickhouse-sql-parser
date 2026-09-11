SELECT
    user_id,
    groupArrayArray(page_visits) AS user_session_page_sequences
FROM website_visits
GROUP BY user_id;
