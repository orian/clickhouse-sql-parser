SELECT id, sequenceNextNode('backward', 'tail')(dt, page, page = 'Basket', page = 'Basket', page = 'Gift') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift
1970-01-01 09:00:03    1   Exit // 기준점, Basket과 불일치

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home // 결과값
1970-01-01 09:00:03    2   Gift // Gift와 일치
1970-01-01 09:00:04    2   Basket // 기준점, Basket과 일치

1970-01-01 09:00:01    3   Gift
1970-01-01 09:00:02    3   Home // 결과값
1970-01-01 09:00:03    3   Gift // 기준점, Gift와 일치
1970-01-01 09:00:04    3   Basket // 기준점, Basket과 일치
