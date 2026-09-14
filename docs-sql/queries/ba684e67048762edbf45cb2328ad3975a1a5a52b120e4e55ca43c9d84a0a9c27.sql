SELECT id, sequenceNextNode('backward', 'last_match')(dt, page, ref = 'ref2', page = 'B') FROM test_flow_basecond GROUP BY id;

                  dt   id   page   ref
 1970-01-01 09:00:01    1   A      ref4
 1970-01-01 09:00:02    1   A      ref3 // 결과
 1970-01-01 09:00:03    1   B      ref2 // 기준점
 1970-01-01 09:00:04    1   B      ref1 // 이 행은 ref 컬럼이 'ref2'와 일치하지 않으므로 기준점이 될 수 없습니다.
