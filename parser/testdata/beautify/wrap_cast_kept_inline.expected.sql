SELECT
  toTypeName(
    tuple('a' AS first, 'b' AS second, 'c' AS third)::Tuple(first String, second String, third String)
  )
FROM
  t
