SET use_variant_default_implementation_for_comparisons = 0;
SELECT * FROM test WHERE v2 == [1,2,3]::Array(UInt32)::Variant(String, UInt64, Array(UInt32));
