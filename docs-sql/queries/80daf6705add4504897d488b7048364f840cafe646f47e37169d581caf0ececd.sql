    CREATE TABLE IF NOT EXISTS uk.uk_price_paid_distributed
    ON CLUSTER cluster_2S_2R
    ENGINE = Distributed('cluster_2S_2R', 'uk', 'uk_price_paid_local', rand());
