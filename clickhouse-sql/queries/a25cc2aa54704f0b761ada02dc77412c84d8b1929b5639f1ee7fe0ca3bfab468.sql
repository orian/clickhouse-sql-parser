
        SELECT {func}('hierarchical_{dictionary_layout}_shards{dictionary_shards}_dictionary', number + 1) FROM numbers(1000000) FORMAT Null;
    