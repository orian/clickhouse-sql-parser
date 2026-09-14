
      insert into alter_select_{engine} (part_id, col_0)
      select toString(number % 5000), 0 from numbers(10000)
      settings
      max_block_size=1,
      max_insert_threads=32,
      min_insert_block_size_rows=1,
      insert_deduplicate=false,
      parts_to_delay_insert=100000,
      parts_to_throw_insert=100000
  