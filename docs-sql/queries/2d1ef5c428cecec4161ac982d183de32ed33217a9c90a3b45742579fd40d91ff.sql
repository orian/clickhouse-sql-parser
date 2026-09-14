SELECT number
FROM numbers()
WHERE number >= 1e15
  AND bitAnd(sipHash64(number), 0xFFFFF) = 0
LIMIT 1;
