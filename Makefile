openapi:
    # Генерация кода для users
	oapi-codegen -package users -o internal/web/users/api.gen.go openapi.yaml
lint:
	golangci-lint run --out-format colored-line-number
