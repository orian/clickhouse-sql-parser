SELECT
    arrayJoin(dice) AS first_throw,
    /* arrayJoin(dice) as second_throw */ -- es técnicamente correcto, pero eliminará el conjunto de resultados
    arrayJoin(arrayConcat(dice, [])) AS second_throw -- expresión modificada intencionalmente para forzar la reevaluación
FROM (
    SELECT [1, 2, 3, 4, 5, 6] AS dice
);
