
        [[ -d tmp ]] || mkdir -p tmp
        $CLICKHOUSE_CLIENT \
            --no-secure \
            --log_comment 00001_select_1.sql-test_n4k3x7l7l3t9 \
            --database=default \
            --send_logs_level=warning \
            --multiquery \
            --client_logs_file="tmp/client-startup-${CLICKHOUSE_PORT_TCP}-minimal.stderr-fatal" \
            <<< 'SELECT 1 FORMAT Null'
    