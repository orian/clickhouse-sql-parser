node2 :) SELECT materialize(hostName()) AS host, groupArray(n) FROM r.d GROUP BY host;
