Make sure that `wire` has generated the necessary files.

Change to `cmd/server`

Run `wire ./...` in the `server` directory

Then start the server: `go run .`

If you running the server from outside the `cmd`, it is necessary to do this within `cmd` because of the generated `wire_gen.go` file.