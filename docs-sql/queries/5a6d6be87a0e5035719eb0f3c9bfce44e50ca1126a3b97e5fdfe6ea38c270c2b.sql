SELECT d, dynamicType(d) FROM test GROUP BY d SETTINGS allow_suspicious_types_in_group_by=1;
