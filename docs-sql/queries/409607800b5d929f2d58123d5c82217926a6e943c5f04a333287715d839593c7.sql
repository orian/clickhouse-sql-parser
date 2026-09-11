SELECT id, sequenceNextNode('forward', 'first_match')(dt, page, page = 'Gift', page = 'Gift', page = 'Home') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift // Базовая точка
1970-01-01 09:00:03    1   Exit // Не совпадает с Home

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home
1970-01-01 09:00:03    2   Gift // Базовая точка
1970-01-01 09:00:04    2   Basket // Не совпадает с Home

1970-01-01 09:00:01    3   Gift // Базовая точка
1970-01-01 09:00:02    3   Home // Совпадает с Home
1970-01-01 09:00:03    3   Gift // Результат
1970-01-01 09:00:04    3   Basket
