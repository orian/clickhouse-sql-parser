SELECT
    toTypeName(toLowCardinality('') AS val) AS source_type,
    toTypeName(toString(val)) AS to_type_result_type,
    toTypeName(CAST(val, 'String')) AS cast_result_type

┌─source_type────────────┬─to_type_result_type────┬─cast_result_type─┐
│ LowCardinality(String) │ LowCardinality(String) │ String           │
└────────────────────────┴────────────────────────┴──────────────────┘

SELECT
    toTypeName(toNullable('') AS val) AS source_type,
    toTypeName(toString(val)) AS to_type_result_type,
    toTypeName(CAST(val, 'String')) AS cast_result_type

┌─source_type──────┬─to_type_result_type─┬─cast_result_type─┐
│ Nullable(String) │ Nullable(String)    │ String           │
└──────────────────┴─────────────────────┴──────────────────┘

SELECT
    toTypeName(toNullable('') AS val) AS source_type,
    toTypeName(toString(val)) AS to_type_result_type,
    toTypeName(CAST(val, 'String')) AS cast_result_type
SETTINGS cast_keep_nullable = 1

┌─source_type──────┬─to_type_result_type─┬─cast_result_type─┐
│ Nullable(String) │ Nullable(String)    │ Nullable(String) │
└──────────────────┴─────────────────────┴──────────────────┘
