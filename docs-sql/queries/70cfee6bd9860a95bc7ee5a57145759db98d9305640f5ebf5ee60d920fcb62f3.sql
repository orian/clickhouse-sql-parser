SELECT toTypeName(variant) AS type_name, [1,2,3]::Array(UInt64)::Variant(UInt64, String, Array(UInt64)) as variant, variantType(variant) as variant_name
