SELECT id, sequenceNextNode('forward', 'head')(dt, page, page = 'Home', page = 'Home', page = 'Gift') FROM test_flow GROUP BY id;

                  dt   id   page
 1970-01-01 09:00:01    1   Home // Базовая точка, совпадение с Home
 1970-01-01 09:00:02    1   Gift // Совпадение с Gift
 1970-01-01 09:00:03    1   Exit // Результат

 1970-01-01 09:00:01    2   Home // Базовая точка, совпадение с Home
 1970-01-01 09:00:02    2   Home // Нет совпадения с Gift
 1970-01-01 09:00:03    2   Gift
 1970-01-01 09:00:04    2   Basket

 1970-01-01 09:00:01    3   Gift // Базовая точка, нет совпадения с Home
 1970-01-01 09:00:02    3   Home
 1970-01-01 09:00:03    3   Gift
 1970-01-01 09:00:04    3   Basket
