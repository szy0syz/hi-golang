# Simple Bank

## Notes

```bash
brew install golang-migrate
mkdir -p db/migration
migrate create -ext sql -dir db/migration -seq init_schema

# 接下来的操作竟然是手动复制up和down的sql语句...😅 
```
