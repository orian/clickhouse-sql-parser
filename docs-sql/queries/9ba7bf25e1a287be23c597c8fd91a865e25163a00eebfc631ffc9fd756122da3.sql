  -- 컬럼 타입이 예상과 일치하는지 확인합니다
  DESCRIBE TABLE events FORMAT Vertical

  -- 데이터를 삽입하고 쿼리해 스키마가 데이터를 올바르게 처리하는지 검증합니다
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
