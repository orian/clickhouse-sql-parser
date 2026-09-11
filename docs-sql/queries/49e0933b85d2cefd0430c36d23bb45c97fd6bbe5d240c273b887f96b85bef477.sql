  CREATE FOREIGN TABLE test (
      column1 bigint  OPTIONS(AggregateFunction 'uniq'),
      column2 integer OPTIONS(AggregateFunction 'anyIf'),
      column3 bigint  OPTIONS(AggregateFunction 'quantiles(0.5, 0.9)')
  ) SERVER clickhouse_srv;
