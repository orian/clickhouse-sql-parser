  -- ClickHouse ネイティブのクエリパラメータ（推奨）
  SET param_user_id = 12345;
  SELECT * FROM users WHERE id = {user_id: UInt64};
