  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- تأكيد سلامة البيانات بعد الكتابة ثم القراءة
  SELECT payload FROM raw_events WHERE id = 1

  -- التحقق من أنه لا يزال بإمكانك تحليل الحقول عند الحاجة بشكل مخصص
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
