  INSERT INTO dynamic_events (id, data) FORMAT JSONEachRow
  {"id": 1, "data": {"event_type": "click", "version": 2, "page": "/home", "button_id": "cta-1"}}
  {"id": 2, "data": {"event_type": "purchase", "version": 1, "item_id": "SKU-99", "amount": 49.99, "currency": "USD"}}

  -- 检查 ClickHouse 检测到的路径及其类型
  SELECT JSONAllPathsWithTypes(data) FROM dynamic_events FORMAT PrettyJSONEachRow

  -- 查询特定路径
  SELECT data.page FROM dynamic_events WHERE data.event_type = 'click'
