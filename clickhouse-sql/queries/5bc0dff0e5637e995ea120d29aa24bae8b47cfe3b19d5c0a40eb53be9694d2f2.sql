
         SELECT key, value
         FROM mingroupby_orderbylimit1_{aggregationscale_small}_tuple
         order by key, value limit 1 by key format Null;
     