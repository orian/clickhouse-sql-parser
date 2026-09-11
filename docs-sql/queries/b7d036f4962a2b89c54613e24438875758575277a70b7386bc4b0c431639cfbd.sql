SELECT id, sequenceNextNode('forward', 'head')(dt, page, page = 'Home', page = 'Home', page = 'Gift') FROM test_flow GROUP BY id;

                  dt   id   page
 1970-01-01 09:00:01    1   Home // 기준점, Home과 일치함
 1970-01-01 09:00:02    1   Gift // Gift와 일치함
 1970-01-01 09:00:03    1   Exit // 결과

 1970-01-01 09:00:01    2   Home // 기준점, Home과 일치함
 1970-01-01 09:00:02    2   Home // Gift와 일치하지 않음
 1970-01-01 09:00:03    2   Gift
 1970-01-01 09:00:04    2   Basket

 1970-01-01 09:00:01    3   Gift // 기준점, Home과 일치하지 않음
 1970-01-01 09:00:02    3   Home
 1970-01-01 09:00:03    3   Gift
 1970-01-01 09:00:04    3   Basket
