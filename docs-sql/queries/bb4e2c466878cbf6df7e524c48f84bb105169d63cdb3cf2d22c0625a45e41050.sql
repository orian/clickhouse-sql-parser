    CREATE TABLE events (
       event_id SERIAL PRIMARY KEY,
       event_name VARCHAR(255) NOT NULL,
       event_type VARCHAR(100),
       event_timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       event_data JSONB,
       user_id INT,
       user_ip INET,
       is_active BOOLEAN DEFAULT TRUE,
       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE users (
       user_id SERIAL PRIMARY KEY,
       name VARCHAR(100),
       country VARCHAR(50),
       platform VARCHAR(50)
    );
