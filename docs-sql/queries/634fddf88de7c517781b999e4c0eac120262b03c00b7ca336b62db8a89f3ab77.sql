        -- resultTable은 clickhouse.default.example_table에 삽입할 Spark 중간 DataFrame입니다.
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
