SELECT id, sequenceNextNode('forward', 'first_match')(dt, page, page = 'Gift', page = 'Gift', page = 'Home') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift // 기준점
1970-01-01 09:00:03    1   Exit // Home과 일치하지 않음

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home
1970-01-01 09:00:03    2   Gift // 기준점
1970-01-01 09:00:04    2   Basket // Home과 일치하지 않음

1970-01-01 09:00:01    3   Gift // 기준점
1970-01-01 09:00:02    3   Home // Home과 일치함
1970-01-01 09:00:03    3   Gift // 결과값
1970-01-01 09:00:04    3   Basket
