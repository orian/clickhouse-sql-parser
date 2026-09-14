SELECT count() FROM
        (SELECT materialize(0) AS x1, materialize(0) AS x2, materialize(0) AS x3, materialize(0) AS x4, materialize(0) AS x5, materialize(0) AS x6, materialize(0) AS x7, materialize(0) AS x8, materialize(0) AS x9, materialize(0) AS x10 FROM zeros(20000000))
    WHERE NOT ignore(or(x1,x2,x3,x4,x5,x6,x7,x8,x9,x10))