CREATE TABLE paimon_mv_dest AS paimon_mv_source
ENGINE = MergeTree()
ORDER BY tuple();
