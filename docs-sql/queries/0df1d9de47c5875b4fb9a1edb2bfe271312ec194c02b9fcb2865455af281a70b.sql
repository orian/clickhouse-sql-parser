SOURCE(EXECUTABLE_POOL(
    command 'while read key; do printf "$key\tData for key $key\n"; done'
    format 'TabSeparated'
    pool_size 10
    max_command_execution_time 10
    implicit_key false
))
