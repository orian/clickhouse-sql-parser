  -- Parâmetros de consulta nativos do ClickHouse (recomendado)
  SET param_user_id = 12345;
  SELECT * FROM users WHERE id = {user_id: UInt64};
