  -- Confirma que los tipos de columna coincidan con lo esperado
  DESCRIBE TABLE events FORMAT Vertical

  -- Inserta y consulta para validar que el esquema procese correctamente tus datos
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
