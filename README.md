Setup
====
```
go mod download github.com/akavel/rsrc1
go get topfrag.org/tfparser
go get -u github.com/asticode/go-astilectron-bundler/...
go install github.com/asticode/go-astilectron-bundler/astilectron-bundler
cd frontend
npm install
npm run build
cd ..
astilectron-bundler
```
