SELECT toDate(CreationDate) AS day, avg(ViewCount) AS view_count FROM stackoverflow.posts WHERE day > '2009-01-01'  GROUP BY day
