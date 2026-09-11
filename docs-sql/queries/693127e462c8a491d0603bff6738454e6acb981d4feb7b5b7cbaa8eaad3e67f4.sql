        -- resultTable は、clickhouse.default.example_table に挿入する Spark の中間 DataFrame です
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
