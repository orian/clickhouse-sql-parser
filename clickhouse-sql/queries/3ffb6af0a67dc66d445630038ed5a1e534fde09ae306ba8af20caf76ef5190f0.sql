
        insert into ht_prefetch select concat('key', toString(number)), concat('value', toString(rand())), rand() % 1000 from numbers(1e7)
    