.PHONY: clean go js

run:
	$(MAKE) clean
	$(MAKE) go
	$(MAKE) js

GO_OUTPUT := go

go:
	@mkdir -p ./go
	@npx openapi-generator-cli generate -g go -i api.yaml -o $(GO_OUTPUT) --git-repo-id="mint-sdk-api/go" --git-user-id="KyuzanInc"

JS_OUTPUT := js/src/apiClient

js:
	@mkdir -p ./js
	@npx openapi-generator-cli generate -g typescript-axios -i api.yaml -o $(JS_OUTPUT)

clean:
	rm -rf \
		$(GO_OUTPUT) \
		$(JS_OUTPUT)
