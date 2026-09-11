
    INSERT INTO text_index_build_src_5000 SELECT arrayStringConcat(arrayMap(x -> toString(rand(number * 30 + x) % 5000), range(30)), ' ') FROM numbers(100000);
  