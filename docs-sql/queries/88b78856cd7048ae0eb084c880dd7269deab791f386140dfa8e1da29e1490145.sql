  -- 샘플 데이터를 삽입하고 추론된 경로를 확인합니다
  INSERT INTO events FORMAT JSONEachRow
  {"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42,"attributes":{"http.status_code":200,"http.method":"GET","user.region":"eu-west","custom_tag":"abc"}}

  SELECT JSONAllPathsWithTypes(attributes)
  FROM events
  FORMAT PrettyJSONEachRow
