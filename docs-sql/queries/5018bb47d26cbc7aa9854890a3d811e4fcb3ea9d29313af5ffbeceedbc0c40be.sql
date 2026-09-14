  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- Confirma que los datos se recuperan intactos
  SELECT payload FROM raw_events WHERE id = 1

  -- Verifica que aún puedes extraer campos de forma ad hoc cuando sea necesario
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
