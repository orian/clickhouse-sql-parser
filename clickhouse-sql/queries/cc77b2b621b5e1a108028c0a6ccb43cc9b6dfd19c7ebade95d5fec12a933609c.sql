
        INSERT INTO text_index_trivial_count
        SELECT number,
               arrayStringConcat(arrayFilter(x -> x != '', [
                   if(number % 2 = 0, 'thecommon', ''),
                   if(number % 1000 = 0, 'rareword', ''),
                   concat('w', toString(number % 5000))]), ' ')
        FROM numbers(20000000)
    