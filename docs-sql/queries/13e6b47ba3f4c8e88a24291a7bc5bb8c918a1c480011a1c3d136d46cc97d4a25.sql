SELECT comment, hex(decrypt('aes-256-cfb8', secret, '12345678910121314151617181920212')) AS plaintext FROM encryption_test
