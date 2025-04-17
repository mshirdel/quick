package config

var _builtinConfig = `
server:
  address: '0.0.0.0:8080'
  read-timeout: 1s
  write-timeout: 3s
  idle-timeout: 2m

http:
  body-limit-size: 1MB
  cors:
    allowed-origins: ['*']
    allowed-headers: []
    allowed-methods: ['GET', 'POST', 'PATCH', 'DELETE']
    allow-credentials: false
    exposed-headers: []
    max-age: 0
  recover:
    stack-size: 4
    disable-stack-all: false
    disable-print-stack: false
database:
  host: localhost
  port: 5432
  user: postgres 
  password: 123456
  dbname: quick 
  connection-timeout: 30s
  connection-lifetime: 30m
  pool-size: 100
  max-idle-connections: 10
  logger:
    slow-threshold: 200ms
`
