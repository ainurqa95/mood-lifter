build:
	    docker build -t mood-lifter:v1 .
migrate:
		goose -dir db/migrations postgres "postgresql://ainur:secret@127.0.0.1:5432/mood_lifter?sslmode=disable" up
