        concat(
          remote_addr, ' ',
          remote_user, ' ',
          '[', formatDateTime(time_local, '%d/%b/%Y:%H:%M:%S %z'), '] ',
          '"', request, '" ',
          toString(status), ' ',
          toString(body_bytes_sent), ' ',
          '"', http_referer, '" ',
          '"', http_user_agent, '" ',
          '"', http_x_forwarded_for, '" ',
          toString(request_time), ' ',
          toString(upstream_response_time), ' ',
          '"', http_host, '"'
        )
