#!/bin/bash

DEPLOY=${1:-cloud-run-lifecycle}
echo $DEPLOY
gcloud beta run deploy --source=. --platform=managed --region=us-central1 --allow-unauthenticated $DEPLOY $@