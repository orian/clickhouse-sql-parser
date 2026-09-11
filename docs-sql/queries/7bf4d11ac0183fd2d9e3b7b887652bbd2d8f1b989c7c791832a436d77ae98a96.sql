SELECT
    user_agent,
    count(DISTINCT user_id)
FROM visits
GROUP BY user_agent
