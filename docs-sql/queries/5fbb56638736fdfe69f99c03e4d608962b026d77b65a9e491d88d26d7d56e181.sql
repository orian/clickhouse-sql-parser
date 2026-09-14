  SELECT
     hostname() AS host,
     count()
     FROM system.mutations WHERE not is_done
     GROUP BY host;
