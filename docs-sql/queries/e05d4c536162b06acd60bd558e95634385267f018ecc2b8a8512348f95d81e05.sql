CREATE TABLE apify_products
(
    url        String,
    title      String,
    price      Float64,
    currency   String,
    scraped_at DateTime DEFAULT now()
)
ENGINE = MergeTree()
ORDER BY (scraped_at, url);
