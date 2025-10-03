curl -X POST http://localhost:8080/data ^
  -H "Content-Type: application/json" ^
  -d "{\"device_id\":\"test\",\"temperature\":25.5,\"humidity\":60}"