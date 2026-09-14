CREATE TABLE uk.uk_price_paid
(
  ...
)
ENGINE = MergeTree
ORDER BY (postcode1, postcode2, addr1, addr2);
