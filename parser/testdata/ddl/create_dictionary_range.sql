CREATE DICTIONARY test.discounts (
    advertiser_id UInt64,
    discount_start_date Date,
    discount_end_date Date,
    amount Float64 DEFAULT 0
)
PRIMARY KEY advertiser_id
SOURCE(CLICKHOUSE(TABLE 'discounts' DB 'test'))
LIFETIME(MIN 1 MAX 1000)
LAYOUT(RANGE_HASHED(range_lookup_strategy 'max'))
RANGE(MIN discount_start_date MAX discount_end_date)
SETTINGS(format_csv_allow_single_quotes = 0);
