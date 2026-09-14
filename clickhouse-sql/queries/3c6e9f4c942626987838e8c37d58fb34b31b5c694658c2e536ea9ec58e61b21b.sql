
    INSERT INTO text_index_build_src_50 SELECT arrayStringConcat(arrayMap(x -> toString(rand(number * 30 + x) % 50), range(30)), ' ') FROM numbers(100000);
  