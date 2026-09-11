
        INSERT INTO polys
        SELECT
            CAST(
                [
                    /* outer ring (closed) */
                    [
                        tuple(cx - r, cy - r),
                        tuple(cx + r, cy - r),
                        tuple(cx + r, cy + r),
                        tuple(cx - r, cy + r),
                        tuple(cx - r, cy - r)
                    ],
                    /* hole 1 (closed, CW) */
                    [
                        tuple(cx - dx - h1, cy - h1),
                        tuple(cx - dx - h1, cy + h1),
                        tuple(cx - dx + h1, cy + h1),
                        tuple(cx - dx + h1, cy - h1),
                        tuple(cx - dx - h1, cy - h1)
                    ],
                    /* hole 2 (closed, CW) */
                    [
                        tuple(cx + dx - h2, cy - h2),
                        tuple(cx + dx - h2, cy + h2),
                        tuple(cx + dx + h2, cy + h2),
                        tuple(cx + dx + h2, cy - h2),
                        tuple(cx + dx - h2, cy - h2)
                    ]
                ] AS Array(Array(Tuple(Int32, Int32)))
            ) AS poly
        FROM
        (
            SELECT
                number % 40 AS col,
                intDiv(number, 40) AS row,
                toInt32(toInt32(col) * 50 - 975) AS cx,
                toInt32(toInt32(row) * 50 - 600) AS cy,
                toInt32(20 + intHash64(number) % 5) AS r,
                toInt32(intDiv(r, 3)) AS dx,
                toInt32(intDiv(r, 3)) AS h1,
                toInt32(intDiv(r, 4)) AS h2
            FROM numbers(1000)
        );
    