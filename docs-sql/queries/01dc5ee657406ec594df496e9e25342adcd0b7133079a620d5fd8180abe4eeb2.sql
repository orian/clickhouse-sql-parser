SELECT id, sequenceNextNode('forward', 'first_match')(dt, page, ref = 'ref3', page = 'A') FROM test_flow_basecond GROUP BY id;

                  dt   id   page   ref
 1970-01-01 09:00:01    1   A      ref4 // この行は、refカラムが 'ref3' と一致しないため、基準点 にはなりません。
 1970-01-01 09:00:02    1   A      ref3 // 基準点
 1970-01-01 09:00:03    1   B      ref2 // 結果
 1970-01-01 09:00:04    1   B      ref1
