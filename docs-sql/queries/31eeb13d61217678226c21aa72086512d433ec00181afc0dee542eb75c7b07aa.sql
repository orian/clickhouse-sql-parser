    CREATE FOREIGN TABLE hits (
        watchid    bigint   OPTIONS(column_name 'WatchID'),
        javaenable smallint OPTIONS(column_name 'JavaEnable'),
        title      text     OPTIONS(column_name 'Title')
    ) SERVER taxi_srv OPTIONS(table_name 'hits');
