       CREATE TEMPORARY VIEW jdbcTable
               USING org.apache.spark.sql.jdbc
               OPTIONS (
                       url "jdbc:ch://localhost:8123/default", 
                       dbtable "schema.tablename",
                       user "username",
                       password "password",
                       driver "com.clickhouse.jdbc.ClickHouseDriver" 
               );
       -- resultTable pode ser criada com df.createTempView ou com Spark SQL
       INSERT INTO TABLE jdbcTable
                    SELECT * FROM resultTable;
                    
