
        CREATE TABLE test_avg_insert (key UInt64, value AggregateFunction(avg, UInt8)) ENGINE = Memory() SETTINGS max_rows_to_keep=1
    