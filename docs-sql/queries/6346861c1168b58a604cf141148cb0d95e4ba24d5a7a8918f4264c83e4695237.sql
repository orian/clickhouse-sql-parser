  SELECT formatReadableSize(total_bytes)
  FROM system.tables
  WHERE name = 'uk_price_paid'
