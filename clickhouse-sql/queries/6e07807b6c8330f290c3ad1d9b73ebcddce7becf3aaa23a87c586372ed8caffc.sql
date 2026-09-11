
        SELECT COUNT(*) AS count
        FROM
        (
            SELECT
                (number % {num_unique_sessions}) AS visitor_id,
                number AS id
            FROM system.numbers
            LIMIT {num_rows}
        ) AS sessions
        ASOF LEFT JOIN
        (
            SELECT
                (number % {num_unique_sessions}) AS visitor_id,
                number AS starting_session_id
            FROM system.numbers
            LIMIT {num_rows}
        ) AS visitors
        ON visitors.visitor_id = sessions.visitor_id AND visitors.starting_session_id < sessions.id
    