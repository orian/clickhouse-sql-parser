CREATE TABLE mytable
(
    x String CODEC(AES_128_GCM_SIV)
)
ENGINE = MergeTree ORDER BY x;
