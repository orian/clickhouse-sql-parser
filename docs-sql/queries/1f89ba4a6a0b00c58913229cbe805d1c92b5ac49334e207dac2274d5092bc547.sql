SELECT count()
FROM paimon_inc
SETTINGS max_consume_snapshots = 2;
