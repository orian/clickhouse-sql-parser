SELECT
        Day,
        UpVotes,
        DownVotes
FROM up_down_votes_per_day
FINAL
ORDER BY Day ASC
LIMIT 10
