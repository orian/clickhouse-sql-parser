
        ${CLICKHOUSE_CURL} -H 'Accept-Encoding: gzip' \
            "${CLICKHOUSE_URL}?enable_http_compression=1" \
            --data-binary "SELECT number FROM numbers(500000) FORMAT TSV" \
            -o /dev/null
    