CREATE FOREIGN TABLE acts (
    user_id    bigint NOT NULL,
    page_views int,
    duration   smallint,
    sign       smallint
) SERVER taxi_srv OPTIONS(
    table_name 'acts'
    engine 'CollapsingMergeTree(sign)'
);
