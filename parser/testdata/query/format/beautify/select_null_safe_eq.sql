-- Origin SQL:
SELECT a <=> b, NOT (c <=> d) FROM t WHERE x <=> y;


-- Beautify SQL:
SELECT
  a <=> b,
  NOT (c <=> d)
FROM
  t
WHERE
  x <=> y;
