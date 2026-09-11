        -- resultTable 是要插入到 clickhouse.default.example_table 的 Spark 中间 DataFrame
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
