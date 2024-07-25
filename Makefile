build: setup depend compile create-html dist deploy cleanup

build-test: setup depend compile create-html

setup: 
	@clear
	@echo "Starting Build Process..."
	@echo "Creating build folder."
	@mkdir -p build
	@echo "build folder created"
	@echo "Copying wasm_exec.js to build folder"
	@cp `go env GOROOT`/misc/wasm/wasm_exec.js build/
	@echo "wasm_exec.js copied to build folder"
	@echo "Copying assets to build folder"
	@cp -r assets build
	@echo "Assets copied to build folder"

create-html:
	@echo "Creating HTML file"
	@echo '<!DOCTYPE html><script src="wasm_exec.js"></script><script>if(!WebAssembly.instantiateStreaming){WebAssembly.instantiateStreaming=async(resp,importObject)=>{const source=await(await resp).arrayBuffer();return await WebAssembly.instantiate(source,importObject);};}const go=new Go();WebAssembly.instantiateStreaming(fetch("game.wasm"),go.importObject).then(result=>{go.run(result.instance);});</script>' > build/index.html
	@echo "HTML file created"

compile:
	@echo "Compiling Game to Wasm"
	@env GOOS=js GOARCH=wasm go build -o build/game.wasm github.com/{profileName}/{gameName}
	@echo "Game Compiled..."

dist: 
	@echo "Creating distribution zip file"
	@zip -r dist.zip build
	@echo "Distribution zip file created"
	
deploy:
	@echo "Uploading to itch.io"
	@butler push dist.zip {profileName}/{gameName}:HTML5 --userversion 0.0.1
	@echo "Upload completed"

cleanup: 
	@echo "Cleaning Up..."
	@rm -rf build
	@rm dist.zip
	@echo "Clean up Completed"
	
status: 
	@butler status {profileName}/{gameName}:HTML5

depend:
	@echo "Updating Dependancies..."
	@go mod tidy
	@echo "Dependancies Updated"

demo: 
	@go run .

test: depend demo

.PHONY: build setup compile create-html dist deploy cleanup status depend demo test