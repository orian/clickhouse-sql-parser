SELECT
    arrayJoin(dice) AS first_throw,
    /* arrayJoin(dice) as second_throw */ -- 기술적으로는 올바르지만, 결과 집합을 소멸시킵니다
    arrayJoin(arrayConcat(dice, [])) AS second_throw -- 재평가를 강제하기 위해 의도적으로 표현식을 변경
FROM (
    SELECT [1, 2, 3, 4, 5, 6] AS dice
);
