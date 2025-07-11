#!/bin/bash

NAME=${1:-cloud-run-lifecycle}
TAG=${2:-0.01}
REGION=asia-northeast1

gcloud artifacts repositories create --repository-format=docker --location=$REGION myrepo
gcloud builds submit --pack=image=$REGION-docker.pkg.dev/$GOOGLE_CLOUD_PROJECT/myrepo/$NAME:$TAG,env=NAME=$NAME
