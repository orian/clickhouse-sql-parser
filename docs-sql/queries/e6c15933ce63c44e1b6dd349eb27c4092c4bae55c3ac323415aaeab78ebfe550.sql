    TRUNCATE TABLE uk.uk_price_paid_local ON CLUSTER cluster_1S_2R;
    INSERT INTO uk.uk_price_paid_local
    SELECT
        toUInt32(price_string) AS price,
        parseDateTimeBestEffortUS(time) AS date,
        splitByChar(' ', postcode)[1] AS postcode1,
        splitByChar(' ', postcode)[2] AS postcode2,
        transform(a, ['T', 'S', 'D', 'F', 'O'], ['terraced', 'semi-detached', 'detached', 'flat', 'other']) AS type,
        b = 'Y' AS is_new,
        transform(c, ['F', 'L', 'U'], ['freehold', 'leasehold', 'unknown']) AS duration,
        addr1,
        addr2,
        street,
        locality,
        town,
        district,
        county
    FROM url(
        'http://prod1.publicdata.landregistry.gov.uk.s3-website-eu-west-1.amazonaws.com/pp-complete.csv',
        'CSV',
        'uuid_string String,
        price_string String,
        time String,
        postcode String,
        a String,
        b String,
        c String,
        addr1 String,
        addr2 String,
        street String,
        locality String,
        town String,
        district String,
        county String,
        d String,
        e String'
        ) SETTINGS max_http_get_redirects=10;
