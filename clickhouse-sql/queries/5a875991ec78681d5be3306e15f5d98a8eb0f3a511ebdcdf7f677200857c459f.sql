
    INSERT INTO text_index_build_src_500000 SELECT arrayStringConcat(arrayMap(x -> toString(rand(number * 30 + x) % 500000), range(30)), ' ') FROM numbers(100000);
  