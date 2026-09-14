SET allow_suspicious_types_in_order_by = 1;
SELECT v2, variantType(v2) AS v2_type FROM test ORDER BY v2;
