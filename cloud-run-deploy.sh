#!/bin/bash

DEPLOY=${1:-cloud-run-lifecycle}
NAME=${2:-cloud-run-lifecycle}
echo $DEPLOY
gcloud beta run deploy --source=. --platform=managed --region=us-central1 --set-env-vars=NAME=$NAME $DEPLOY
# gcloud beta run deploy --source=. --platform=managed --region=us-central1 --allow-unauthenticated $DEPLOY $@
