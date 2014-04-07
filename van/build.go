if [ ! -h src/github.com/javouhey/seneca ]; then
        mkdir -p src/github.com/javouhey/
        ln -s ../../.. src/github.com/javouhey/seneca // HL
fi
export GOPATH=${PWD}

VERSION=$(cat ./VERSION) // HL
GITSHA=$(git rev-parse HEAD) // HL
LDFLAGS='-X main.GitSHA "'$GITSHA'" -X main.Version "'$VERSION'" -w' // HL
go install -ldflags "$LDFLAGS" github.com/javouhey/seneca
