.PHONY: gomodgen
gomodgen:
	GO111MODULE=on go mod init

.PHONY: deploy
deploy:
	gcloud functions deploy hello --entry-point Hello --runtime go111 --trigger-http

.PHONY: delete
delete:
	gcloud functions delete hello --entry-point Hello --runtime go111 --trigger-http