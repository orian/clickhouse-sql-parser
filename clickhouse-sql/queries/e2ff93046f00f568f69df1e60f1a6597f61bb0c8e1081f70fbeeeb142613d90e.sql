
      INSERT INTO tab
      SELECT
          number,
          concat('prefix', toString(number), 'suffix')
      FROM numbers({table_size})
  