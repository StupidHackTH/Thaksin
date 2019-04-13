build:
	gcloud builds submit --tag asia.gcr.io/phoomparin/thaksin
	gcloud beta run deploy thaksin --image asia.gcr.io/phoomparin/thaksin --project phoomparin

build-local:
	docker build -t asia.gcr.io/phoomparin/thaksin:latest .
	docker push asia.gcr.io/phoomparin/thaksin:latest
	gcloud beta run deploy thaksin --image asia.gcr.io/phoomparin/thaksin --project phoomparin
