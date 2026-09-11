SELECT
    user_agent,
    sum(pages_visited)
FROM visits
GROUP BY user_agent
