  -- Убедитесь, что типы столбцов соответствуют ожидаемым
  DESCRIBE TABLE events FORMAT Vertical

  -- Выполните вставку и запрос, чтобы проверить, что схема корректно обрабатывает ваши данные
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
