SELECT Hour, uniqMerge(UniqueUsers) AS UniqueUsers
FROM unique_visitors_per_hour
GROUP BY Hour
ORDER BY Hour DESC
