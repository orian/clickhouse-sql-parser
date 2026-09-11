SELECT count() FROM posts WHERE has(keywords, 'clickhouse'); -- 全表扫描，速度较慢——需检查每篇文章中的每个关键词
