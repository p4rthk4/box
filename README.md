# Box SMTP
Box SMTP - Simple Mail Transfer Protocol Server and Client

## Concept
Box SMTP mainly work in SMTP Server Wrapper, It make one server or more than one server (two, three, etc...), Any server recive email and it will be store into database or add into queqe and send/forward to main centrel server.

## TODO
- Enhanced status codes in smtp
- send mail to multiple rcpts
- `exp=` support in  in spf

## Config File

### Config file location
- `/etc/box.yaml`
- `/etc/box.yml`
- `./config.yaml`
- `./config.yml`

### Config file structure
```yaml
name: Box - SMTP
hostname: localhost
max_clients: 5

port: 25

tls:
  starttls: true
  key: key.pem
  cert: cert.pem

spf_check: true
message_size: 1024000
max_recipients: -1
check_mailbox: false

amqp_conf:
  host: localhost
  port: 5672
  username:
  password:
  queue: box-receiver-1
  status_queue: box-status-1

client:
  hostname: localhost
  worker: 2
  log_dir: ./log
  log_file: email.log

  amqp_conf:
    host: localhost
    port: 5672
    username:
    password:
    queue: box-sender-1
    status_queue: box-status-1

log_dir: ./log
log_file: email.log
dev: false
```
