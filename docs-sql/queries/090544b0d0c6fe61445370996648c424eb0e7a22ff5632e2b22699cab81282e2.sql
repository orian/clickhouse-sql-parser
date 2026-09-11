-- Using multiple arrayJoin with the same expression may not produce the expected result due to optimizations.
-- For these cases, consider modifying the repeated array expression with extra operations that do not affect join result.
-- e.g. arrayJoin(arraySort(arr)), arrayJoin(arrayConcat(arr, []))

SELECT
    arrayJoin(dice) as first_throw,
    /* arrayJoin(dice) as second_throw */ -- is technically correct, but will annihilate result set
    arrayJoin(arrayConcat(dice, [])) as second_throw -- intentionally changed expression to force re-evaluation
FROM (
    SELECT [1, 2, 3, 4, 5, 6] as dice
);
