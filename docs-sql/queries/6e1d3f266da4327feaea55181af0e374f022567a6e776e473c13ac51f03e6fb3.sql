       CREATE TABLE db1.table1 on cluster 'cluster_2S_1R'
       (
           `id` UInt64,
           `column1` String
       )
       ENGINE = MergeTree
       ORDER BY column1
