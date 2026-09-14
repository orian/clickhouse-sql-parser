  -- 确认列类型符合预期
  DESCRIBE TABLE events FORMAT Vertical

  -- 通过插入和查询来验证 schema 能否处理你的数据
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
