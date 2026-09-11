SELECT count() FROM posts WHERE has(keywords, 'clickhouse'); -- 느린 전체 테이블 스캔 - 모든 게시물의 모든 키워드를 확인
