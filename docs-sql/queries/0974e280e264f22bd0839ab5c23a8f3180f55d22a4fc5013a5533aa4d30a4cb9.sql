SELECT arrayMap(negate, [1, 2, 3]);            -- [-1, -2, -3]
SELECT arrayMap(x -> negate(x), [1, 2, 3]);    -- [-1, -2, -3]

SELECT arrayMap(plus, [1, 2, 3], [10, 20, 30]);            -- [11, 22, 33]
SELECT arrayMap((x, y) -> plus(x, y), [1, 2, 3], [10, 20, 30]); -- [11, 22, 33]

SELECT arrayFilter(isNotNull, [1, NULL, 3, NULL, 5]);            -- [1, 3, 5]
SELECT arrayFilter(x -> isNotNull(x), [1, NULL, 3, NULL, 5]);    -- [1, 3, 5]

SELECT arrayFold(plus, [1, 2, 3, 4, 5], toUInt64(0));                      -- 15
SELECT arrayFold((acc, x) -> plus(acc, x), [1, 2, 3, 4, 5], toUInt64(0));  -- 15
