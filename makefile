local-akil:
	DB_HOST=localhost DB_PORT=5432 DB_NAME=gomen DB_USER=user DB_PASS=password DB_SSL=disable TIME_ZONE=Asia/Jakarta go run main.go  api --host=0.0.0.0 --port=1712