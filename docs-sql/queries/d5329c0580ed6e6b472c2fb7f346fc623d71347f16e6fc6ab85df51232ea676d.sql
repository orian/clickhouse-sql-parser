        -- resultTable — это промежуточный df Spark, который нужно вставить в clickhouse.default.example_table
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
