
        create table distributed_table_2 (id int) engine = Distributed(test_cluster_two_shard_three_replicas_localhost, currentDatabase(), local_table_2)
    