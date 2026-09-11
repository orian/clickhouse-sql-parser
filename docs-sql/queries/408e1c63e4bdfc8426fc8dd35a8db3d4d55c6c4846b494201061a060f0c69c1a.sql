SELECT
    arrayJoin(dice) AS first_throw,
    /* arrayJoin(dice) as second_throw */ -- технически корректно, но уничтожит результирующий набор
    arrayJoin(arrayConcat(dice, [])) AS second_throw -- выражение намеренно изменено для принудительного повторного вычисления
FROM (
    SELECT [1, 2, 3, 4, 5, 6] AS dice
);
