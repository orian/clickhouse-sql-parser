SELECT
    arrayJoin(dice) AS first_throw,
    /* arrayJoin(dice) as second_throw */ -- 技術的には正しいが、result setを消滅させてしまう
    arrayJoin(arrayConcat(dice, [])) AS second_throw -- 再評価を強制するために意図的に式を変更
FROM (
    SELECT [1, 2, 3, 4, 5, 6] AS dice
);
