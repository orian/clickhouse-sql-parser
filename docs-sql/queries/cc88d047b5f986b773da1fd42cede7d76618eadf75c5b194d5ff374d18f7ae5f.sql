CREATE TABLE raw_page_views
(
    page_id UInt32,
    page_name String,
    response_time_ms UInt32,  -- Tiempo de respuesta de la página en milisegundos
    viewed_at DateTime DEFAULT now()
)
ENGINE = MergeTree()
ORDER BY (page_id, viewed_at);
