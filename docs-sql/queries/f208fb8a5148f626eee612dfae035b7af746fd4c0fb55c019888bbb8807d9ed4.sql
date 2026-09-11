SELECT count() FROM numbers(10) WHERE EXISTS(SELECT number FROM numbers(10) WHERE number > 11);
