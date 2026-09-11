SET use_variant_as_common_type = 1;
SELECT toTypeName(array(range(number), number, 'str_' || toString(number))) as array_of_variants_type from numbers(1);
SELECT array(range(number), number, 'str_' || toString(number)) as array_of_variants FROM numbers(3);
