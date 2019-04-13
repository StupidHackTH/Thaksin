main: build-local

remote: build-remote deploy
local: build-local deploy

deploy:
	gcloud beta run deploy thaksin --image asia.gcr.io/phoomparin/thaksin --project phoomparin

build-remote:
	gcloud builds submit --config cloudbuild.yml .

build-local:
	docker build -t asia.gcr.io/phoomparin/thaksin:latest .
	docker push asia.gcr.io/phoomparin/thaksin:latest
