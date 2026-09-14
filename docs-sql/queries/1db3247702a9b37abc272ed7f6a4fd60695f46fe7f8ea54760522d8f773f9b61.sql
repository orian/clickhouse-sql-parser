       CREATE TEMPORARY VIEW jdbcTable
               USING org.apache.spark.sql.jdbc
               OPTIONS (
                       url "jdbc:ch://localhost:8123/default", 
                       dbtable "schema.tablename",
                       user "username",
                       password "password",
                       driver "com.clickhouse.jdbc.ClickHouseDriver" 
               );
       -- resultTable는 df.createTempView 또는 Spark SQL로 생성할 수 있습니다
       INSERT INTO TABLE jdbcTable
                    SELECT * FROM resultTable;
                    
