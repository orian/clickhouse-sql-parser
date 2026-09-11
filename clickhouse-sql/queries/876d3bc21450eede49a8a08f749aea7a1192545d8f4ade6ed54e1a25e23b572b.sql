
        INSERT INTO test_pk SELECT toDateTime('2001-01-01') + number, number % 16 FROM numbers_mt(1000000);
    