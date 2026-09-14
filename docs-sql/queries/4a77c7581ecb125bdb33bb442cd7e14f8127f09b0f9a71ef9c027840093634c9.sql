  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- Vérifiez que les données sont bien restituées à l'identique
  SELECT payload FROM raw_events WHERE id = 1

  -- Vérifiez que vous pouvez toujours extraire des champs à la demande si nécessaire
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
