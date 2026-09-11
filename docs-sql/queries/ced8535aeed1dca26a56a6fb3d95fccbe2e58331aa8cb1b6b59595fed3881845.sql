CREATE TABLE mt_lc (s LowCardinality(String)) ENGINE = MergeTree ORDER BY tuple()
SETTINGS min_bytes_for_wide_part = 0;

INSERT INTO mt_lc SELECT toString(number % 1000) FROM numbers(1000000);

SELECT substream, codec_block_counts
FROM mergeTreeCodecBlockCounts(currentDatabase(), mt_lc)
WHERE column = 's'
ORDER BY substream;
