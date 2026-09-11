
        INSERT INTO topk_join_l SELECT number, sipHash64(number), repeat('a', 16) FROM numbers(5000000)
    