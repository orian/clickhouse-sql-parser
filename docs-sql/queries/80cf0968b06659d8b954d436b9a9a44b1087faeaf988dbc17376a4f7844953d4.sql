INSERT INTO users (name, country, platform)
SELECT
    first_names[first_idx] || ' ' || last_names[last_idx] AS name,
    CASE
        WHEN random() < 0.25 THEN 'India'
        WHEN random() < 0.5 THEN 'USA'
        WHEN random() < 0.7 THEN 'Germany'
        WHEN random() < 0.85 THEN 'China'
        ELSE 'Other'
    END AS country,
    CASE
        WHEN random() < 0.2 THEN 'iOS'
        WHEN random() < 0.4 THEN 'Android'
        WHEN random() < 0.6 THEN 'Web'
        WHEN random() < 0.75 THEN 'Windows'
        WHEN random() < 0.9 THEN 'MacOS'
        ELSE 'Linux'
    END AS platform
FROM
    generate_series(1, 1000) AS seq
    CROSS JOIN LATERAL (
        SELECT
            array['Alice', 'Bob', 'Charlie', 'Diana', 'Eve', 'Frank', 'Grace', 'Hank', 'Ivy', 'Jack', 'Liam', 'Olivia', 'Noah', 'Emma', 'Sophia', 'Benjamin', 'Isabella', 'Lucas', 'Mia', 'Amelia', 'Aarav', 'Riya', 'Arjun', 'Ananya', 'Wei', 'Li', 'Huan', 'Mei', 'Hans', 'Klaus', 'Greta', 'Sofia'] AS first_names,
            array['Smith', 'Johnson', 'Williams', 'Brown', 'Jones', 'Garcia', 'Miller', 'Davis', 'Martinez', 'Taylor', 'Anderson', 'Thomas', 'Jackson', 'White', 'Harris', 'Martin', 'Thompson', 'Moore', 'Lee', 'Perez', 'Sharma', 'Patel', 'Gupta', 'Reddy', 'Zhang', 'Wang', 'Chen', 'Liu', 'Schmidt', 'Müller', 'Weber', 'Fischer'] AS last_names,
            1 + (seq % 32) AS first_idx,
            1 + ((seq / 32)::int % 32) AS last_idx
    ) AS names;
