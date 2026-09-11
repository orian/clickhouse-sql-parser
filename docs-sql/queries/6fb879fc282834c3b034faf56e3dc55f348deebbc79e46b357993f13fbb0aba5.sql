SELECT
    arrayJoin(dice) AS first_throw,
    /* arrayJoin(dice) as second_throw */ -- é tecnicamente correto, mas aniquilará o result set
    arrayJoin(arrayConcat(dice, [])) AS second_throw -- expressão intencionalmente alterada para forçar reavaliação
FROM (
    SELECT [1, 2, 3, 4, 5, 6] AS dice
);
