
        ${CLICKHOUSE_CURL} "${CLICKHOUSE_URL}" \
            --data-binary "SELECT number FROM numbers(500000) FORMAT TSV" \
            -o /dev/null
    