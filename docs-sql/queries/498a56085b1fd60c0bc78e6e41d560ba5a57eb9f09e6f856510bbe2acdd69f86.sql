  -- カラム型が想定どおりであることを確認
  DESCRIBE TABLE events FORMAT Vertical

  -- スキーマがデータを正しく扱えることを確認するため、insert とクエリを実行
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
