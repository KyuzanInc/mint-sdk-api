.PHONY: clean go js

run:
	$(MAKE) clean
	$(MAKE) go
	$(MAKE) js

go:
	@mkdir -p ./go
	@npx openapi-generator-cli generate -g go -i api.yaml -o go

js:
	@mkdir -p ./js
	@npx openapi-generator-cli generate -g typescript-axios -i api.yaml -o js/src/apiClient

clean:
	rm -rf \
		go \
		js
