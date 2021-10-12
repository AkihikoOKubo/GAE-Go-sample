###############
## AppEngine ##
###############
test: ## Exec ginkgo
	go test `go list ./src/... | grep -v src/handler`

deploy:
	gcloud app deploy src/main/app_dev.yaml --promote

#############
## mockgen ##
#############
# repository
gen-repository-mock: ## generate repository mock
	mockgen -source ./src/domain/repository/firebase.go -destination ./src/domain/repository/mock/firebase.go
	mockgen -source ./src/domain/repository/other.go -destination ./src/domain/repository/mock/other.go
