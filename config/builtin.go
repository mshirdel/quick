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
`
