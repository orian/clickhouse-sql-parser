CREATE TABLE mt (a UInt64) ENGINE = MergeTree ORDER BY a
SETTINGS min_bytes_for_wide_part = 0;

INSERT INTO mt SELECT number FROM numbers(100000);

SELECT column, substream, codec_block_counts
FROM mergeTreeCodecBlockCounts(currentDatabase(), mt);
