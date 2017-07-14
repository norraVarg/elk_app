# elk_app
The purpose of this project is confidential.
## To run
```bash
go run src/main.go
```
## To test
```bash
ab -n 10 -c 3 "http://127.0.0.1:8080/player/1"
ab -n 10 -c 3 "http://127.0.0.1:8080/player/2"
curl "http://127.0.0.1:8080/statistic"
```
