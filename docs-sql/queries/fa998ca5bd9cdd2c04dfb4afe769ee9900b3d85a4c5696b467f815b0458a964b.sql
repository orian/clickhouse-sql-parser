SELECT _table, walkover, count()
FROM merge('atp_matches*')
GROUP BY ALL
ORDER BY _table;
