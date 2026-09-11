
         SELECT key, min(value)
         FROM mingroupby_orderbylimit1_{aggregationscale_big}_tuple
         group by key format Null;
    