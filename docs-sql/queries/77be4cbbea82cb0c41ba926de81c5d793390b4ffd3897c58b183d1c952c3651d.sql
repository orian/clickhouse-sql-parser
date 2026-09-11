ALTER TABLE otel_logs
ADD COLUMN LogAttributeItems Array(String)
ALIAS arrayMap(
  (arr) -> concat(arr.1, '=', arr.2),
  LogAttributes::Array(Tuple(String, String))
)
