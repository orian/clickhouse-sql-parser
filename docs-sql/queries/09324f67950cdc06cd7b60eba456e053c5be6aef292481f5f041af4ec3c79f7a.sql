  -- تأكد من أن أنواع الأعمدة تطابق المتوقع
  DESCRIBE TABLE events FORMAT Vertical

  -- أدرِج بيانات واستعلم للتحقق من أن المخطط يتعامل مع بياناتك
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
