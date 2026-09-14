        -- resultTable é o df intermediário do Spark que queremos inserir em clickhouse.default.example_table
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
