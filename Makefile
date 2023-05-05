APP := ./app
MAIN := ./cmd

serve:
	@reflex --start-service -r '\.go$$' -R '^docs/*' make restart

before:
	@swag init --generalInfo routes/routes.go

build:
	@go build -o $(APP) $(MAIN)

start:
	@$(APP)

restart: before build start

.PHONY: start serve restart before sleep # let's go to reserve rules names