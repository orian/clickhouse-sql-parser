
        SELECT count() FROM numbers(1000000)
        WHERE NOT ignore(pointInPolygon(
            (number % 997 / 997. - 0.5, (number * 7 % 991) / 991. - 0.5),
            arrayMap(x -> (cos(x / 8000. * 2 * pi()), sin(x / 8000. * 2 * pi())), range(8000))))
        SETTINGS max_block_size = 64, validate_polygons = 0;
    