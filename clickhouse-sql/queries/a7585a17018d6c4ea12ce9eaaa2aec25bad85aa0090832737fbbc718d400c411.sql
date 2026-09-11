
        SELECT count() FROM numbers(100000000) WHERE NOT ignore(intHash64(number))
    