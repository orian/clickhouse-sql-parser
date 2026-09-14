
        CREATE TABLE 
            window_test(id Int64, value Int64, partition Int64, msg String) 
        Engine=MergeTree 
        ORDER BY id
    