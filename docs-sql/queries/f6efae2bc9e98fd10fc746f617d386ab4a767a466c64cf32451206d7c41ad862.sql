    kcat -C \
      -b <host>:<port> \
      -t github_out \
      -X security.protocol=sasl_ssl \
      -X sasl.mechanisms=PLAIN \
      -X sasl.username=<username> \
      -X sasl.password=<password> \
      -e -q |
    wc -l
