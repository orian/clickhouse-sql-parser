       clickhouse :) CREATE TABLE repl_table ON CLUSTER cluster_1S_2R
                   (
                       id UInt64,
                       column1 Date,
                       column2 String
                   )
                   ENGINE = ReplicatedMergeTree('/clickhouse/tables/{shard}/default/repl_table', '{replica}' )
                   ORDER BY (id);
