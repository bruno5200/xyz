# Go parameters
GOCMD=go
APP=main.go
GORUN=$(GOCMD) run $(APP) #for clean-architecture structure
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# muestra el log de github (no añadir | head -n 10 | tail -n 5 | sed 's/^/  /')
log:
	git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative

count:
	git rev-list --count HEAD

st:
	git status -sb --untracked-file

ust:
	git restore --staged .

push:
	git push && git rebase dev cert && git push && git rebase cert pre && git push && git switch dev

pull:
	git pull --autostash --rebase

build:
	CGO_ENABLED=0 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_NAME) -v $(APP)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_UNIX) -v $(APP)

build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_MAC) -v $(APP)

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_WIN) -v $(APP)

test:
	$(GOTEST) -v -shuffle=on -count=1 -race -timeout=10m ./... -coverprofile=coverage.out

mod:
	$(GOMOD) tidy

run-bin:
	./$(BINARY_NAME)

#limpia los binarios compilados
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_WIN)

# ejecuta el main.go
run:
	$(GORUN)

dscan:
	docker run --rm --name=osv-scanner -v ${PWD}:/src ghcr.io/google/osv-scanner -L /src/go.mod
