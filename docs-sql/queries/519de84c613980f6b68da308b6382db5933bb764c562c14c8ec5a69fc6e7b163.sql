SELECT UserId, count() AS c FROM badges GROUP BY UserId ORDER BY c DESC LIMIT 5
