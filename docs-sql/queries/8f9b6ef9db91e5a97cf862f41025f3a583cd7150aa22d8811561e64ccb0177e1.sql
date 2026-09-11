       CREATE TEMPORARY VIEW jdbcTable
               USING org.apache.spark.sql.jdbc
               OPTIONS (
                       url "jdbc:ch://localhost:8123/default", 
                       dbtable "schema.tablename",
                       user "username",
                       password "password",
                       driver "com.clickhouse.jdbc.ClickHouseDriver" 
               );
               
       SELECT * FROM jdbcTable;
