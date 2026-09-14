  -- ClickHouse 原生查询参数（推荐）
  SET param_user_id = 12345;
  SELECT * FROM users WHERE id = {user_id: UInt64};
