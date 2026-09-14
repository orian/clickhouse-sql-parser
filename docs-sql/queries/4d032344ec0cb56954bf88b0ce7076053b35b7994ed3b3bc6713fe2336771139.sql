SELECT id, sequenceNextNode('backward', 'last_match')(dt, page, ref = 'ref2', page = 'B') FROM test_flow_basecond GROUP BY id;

                  dt   id   page   ref
 1970-01-01 09:00:01    1   A      ref4
 1970-01-01 09:00:02    1   A      ref3 // 结果
 1970-01-01 09:00:03    1   B      ref2 // 基准点
 1970-01-01 09:00:04    1   B      ref1 // 此行不能作为基准点，因为 ref 列与 'ref2' 不匹配。
