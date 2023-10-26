set dotenv-load

default: list

# lists available tasks
@list:
    just --list

# init project
init:
    # TODO: figure out the command to bring deps

# start the server
dev:
    watchexec -e go -c clear -r go run ./cmd/api

# open the project in the browser
open:
    open "$PROJECT_URL" -a "Google Chrome Canary"

# start a console
# console:
#     rails console

# run tests
test:
    go test

# Open the DB
# db:
#     pgcli $DATABASE_URL

docs:
    tmux new-window -n docs 'python3 -m http.server -b 0.0.0.0 -d ~/Downloads/lets-go-further/html'
