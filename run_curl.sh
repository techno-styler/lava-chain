curl -X POST http://localhost:3334 -H "lava-force-cache-refresh: true" -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","id":"1","method":"status","params":[]}'