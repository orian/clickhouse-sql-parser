SELECT
    name,
    loading_error_message
FROM system.user_defined_functions
WHERE load_status = 'Failed';
