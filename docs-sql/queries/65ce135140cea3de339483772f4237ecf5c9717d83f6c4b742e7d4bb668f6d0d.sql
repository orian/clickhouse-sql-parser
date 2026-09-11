WITH 'all_12_25_7_4' AS lhs, 'all_7_100_10_20' AS rhs
SELECT isMergeTreePartCoveredBy(rhs, lhs), isMergeTreePartCoveredBy(lhs, rhs);
