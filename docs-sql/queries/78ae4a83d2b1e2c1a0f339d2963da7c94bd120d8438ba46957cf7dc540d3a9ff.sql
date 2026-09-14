  SELECT
      table,
      count() AS parts
  FROM system.parts
  WHERE (database = 'uk') AND (table IN ['uk_price_paid_simple', 'uk_price_paid_simple_partitioned']) AND active
  GROUP BY table;

