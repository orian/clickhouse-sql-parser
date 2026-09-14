SELECT column, sumMap(codec_block_counts)
FROM mergeTreeCodecBlockCounts(currentDatabase(), mt)
GROUP BY column;
