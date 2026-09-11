-- See shardNum() example above which also demonstrates shardCount()
CREATE TABLE shard_count_example (dummy UInt8)
ENGINE=Distributed(test_cluster_two_shards_localhost, system, one, dummy);
SELECT shardCount() FROM shard_count_example;
