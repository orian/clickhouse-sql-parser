  -- Confirme se os tipos das colunas correspondem ao esperado
  DESCRIBE TABLE events FORMAT Vertical

  -- Insira dados e consulte para validar se o esquema lida com eles
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
