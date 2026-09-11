        -- resultTable هو DataFrame الوسيط في Spark الذي نريد إدراجه في clickhouse.default.example_table
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
