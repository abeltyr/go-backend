setup-air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

dev:
	clear && ./bin/air

restartBuild:
	rm -rf build && mkdir build
	
buildGo:
	GOOS=linux GOARCH=amd64  go build -a -ldflags="-s -w" -o build/main main.go 
	
compress:
	zip -j build/main.zip build/main

glg:
	echo "install the gqlgen dev" && go get -d github.com/99designs/gqlgen && echo "generate the graphql" && go run github.com/99designs/gqlgen generate

mainBuild:
	go build -o main main.go
