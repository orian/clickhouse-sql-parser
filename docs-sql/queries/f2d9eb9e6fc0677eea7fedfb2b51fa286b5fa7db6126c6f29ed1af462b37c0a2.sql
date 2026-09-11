    CREATE TABLE db1.dist_table (
        id UInt64,
        column1 String
    )
    ENGINE = Distributed(cluster_2S_1R,db1,table1)
