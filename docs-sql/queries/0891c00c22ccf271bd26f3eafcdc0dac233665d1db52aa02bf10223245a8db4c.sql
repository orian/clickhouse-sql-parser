        -- resultTable es el DataFrame intermedio de Spark que queremos insertar en clickhouse.default.example_table
       INSERT INTO TABLE clickhouse.default.example_table
                    SELECT * FROM resultTable;
                    
