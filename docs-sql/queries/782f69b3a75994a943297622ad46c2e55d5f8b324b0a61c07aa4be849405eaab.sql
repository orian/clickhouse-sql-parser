SELECT
    name,
    load_status,
    type,
    command,
    return_type,
    argument_types
FROM system.user_defined_functions
FORMAT Vertical;
