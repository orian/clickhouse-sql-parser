SET use_variant_as_common_type = 1;
SELECT if(number % 2, number, range(number)) as variant FROM numbers(5);
