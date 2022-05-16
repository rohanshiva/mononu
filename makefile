install: 
	cd reactapp && npm install
	cd server && go mod tidy

build: 
	cd reactapp && npm run build 

run: 
	cd server && go run main.go
