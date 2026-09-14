SELECT d, dynamicType(d) FROM test ORDER BY d SETTINGS allow_suspicious_types_in_order_by=1;
