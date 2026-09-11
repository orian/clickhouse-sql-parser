        -- resultTable est le DataFrame Spark intermédiaire que nous voulons insérer dans clickhouse.default.example_table
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
