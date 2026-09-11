CREATE TABLE users
(
  username String,
  firstname String,
  lastname String
)
ENGINE = MergeTree
ORDER BY username
