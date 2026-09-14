SELECT
    arrayJoin(dice) AS first_throw,
    /* arrayJoin(dice) as second_throw */ -- 语法上正确，但会消除结果集
    arrayJoin(arrayConcat(dice, [])) AS second_throw -- 故意修改表达式以强制重新求值
FROM (
    SELECT [1, 2, 3, 4, 5, 6] AS dice
);
