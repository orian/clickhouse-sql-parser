
         SELECT key, min(value)
         FROM mingroupby_orderbylimit1_{aggregationscale_big}_key_value
         group by key format Null;
    