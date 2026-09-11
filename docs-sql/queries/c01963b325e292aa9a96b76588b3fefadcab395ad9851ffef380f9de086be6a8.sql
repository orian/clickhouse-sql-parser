CREATE TABLE example (
    user_id UInt64,
    avg_session_length AggregateFunction(avg, UInt32)
);
