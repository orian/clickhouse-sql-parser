dbxquery query="SELECT pickup_ntaname, count(*) AS count
FROM default.trips GROUP BY pickup_ntaname
ORDER BY count DESC LIMIT 10;" connection="chc"
