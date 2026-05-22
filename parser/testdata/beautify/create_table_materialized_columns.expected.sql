CREATE TABLE persons (
  `id` UUID,
  `properties` String,
  `pmat_email` String MATERIALIZED replaceRegexpAll(
    JSONExtractRaw(properties, 'email'),
    concat('^[', regexpQuoteMeta('"'), ']*|[', regexpQuoteMeta('"'), ']*$'),
    ''
  ) COMMENT 'column_materializer::email',
  `pmat_short` String MATERIALIZED foo(properties),
  INDEX idx_email pmat_email TYPE minmax GRANULARITY 1
)
ENGINE = MergeTree
ORDER BY id
