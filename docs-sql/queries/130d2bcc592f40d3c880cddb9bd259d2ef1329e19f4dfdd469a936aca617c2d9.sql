  -- Vérifier que les types de colonnes correspondent aux attentes
  DESCRIBE TABLE events FORMAT Vertical

  -- Insérer des données et exécuter une requête pour valider que le schéma les prend bien en charge
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

  SELECT service, level, duration_ms FROM events WHERE service = 'api'
