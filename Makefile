main: build-local

PROJECT=phoomparin
IMAGE=asia.gcr.io/${PROJECT}/thaksin
TAG=latest

remote: build-remote deploy
local: build-local deploy

deploy:
	@echo --- Deploying Image: ${IMAGE} ---
	gcloud beta run deploy thaksin --image ${IMAGE} --project ${PROJECT}

build-remote:
	@echo --- Building Image Remotely: ${IMAGE} ---
	gcloud builds submit --config cloudbuild.yml .

build-local:
	@echo --- Building Image Locally: ${IMAGE} ---
	docker build -t ${IMAGE}:${TAG} .
	docker push ${IMAGE}:${TAG}
