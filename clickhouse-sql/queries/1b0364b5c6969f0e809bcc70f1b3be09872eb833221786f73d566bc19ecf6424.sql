
        CREATE TABLE {database}.wr_planned
        (
            l_orderkey Int32, l_partkey Int32, l_suppkey Int32, l_linenumber Int32,
            l_quantity Float64, l_extendedprice Float64, l_discount Float64, l_tax Float64,
            l_returnflag String, l_linestatus String,
            l_shipdate Date, l_commitdate Date, l_receiptdate Date,
            l_shipinstruct String, l_shipmode String, l_comment String
        )
        ENGINE = IcebergLocal((SELECT value FROM system.server_settings WHERE name = 'user_files_path') || '/iceberg_suite_write_wr_planned_' || toString(toUnixTimestamp64Micro(now64(6))) || '/', 'Parquet')
    