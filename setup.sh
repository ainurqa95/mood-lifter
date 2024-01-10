//bin
go install github.com/pressly/goose/v3/cmd/goose@latest

goose -dir db/migrations postgres "postgresql://default:secret@127.0.0.1:5432/mood_lifter?sslmode=disable" up

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://default:secret@127.0.0.1:5432/mood_lifter?sslmode=disable