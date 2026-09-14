SELECT count() FROM posts WHERE has(keywords, 'clickhouse'); -- varredura lenta da tabela inteira - verifica cada keyword em cada post
