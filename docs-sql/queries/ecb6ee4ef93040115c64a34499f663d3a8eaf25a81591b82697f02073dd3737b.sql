       CREATE TEMPORARY VIEW jdbcTable
               USING org.apache.spark.sql.jdbc
               OPTIONS (
                       url "jdbc:ch://localhost:8123/default", 
                       dbtable "schema.tablename",
                       user "username",
                       password "password",
                       driver "com.clickhouse.jdbc.ClickHouseDriver" 
               );
       -- resultTable 可以通过 df.createTempView 或 Spark SQL 创建
       INSERT INTO TABLE jdbcTable
                    SELECT * FROM resultTable;
                    
