SELECT count() FROM
        (SELECT materialize(1) AS x1, materialize(1) AS x2, materialize(1) AS x3, materialize(1) AS x4, materialize(1) AS x5, materialize(1) AS x6, materialize(1) AS x7, materialize(1) AS x8, materialize(1) AS x9, materialize(1) AS x10 FROM zeros(20000000))
    WHERE NOT ignore(and(x1,x2,x3,x4,x5,x6,x7,x8,x9,x10))