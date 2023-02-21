#!/bin/bash

DEPLOY=${1:-cloud-run-lifecycle}
NAME=${2:-cloud-run-lifecycle}
REGION=${3:-asia-northeast1}
echo "Deploying $DEPLOY" "..."
CMD="gcloud run deploy --source=. --region=$REGION --platform=managed --set-env-vars=NAME=$NAME,SLACK_API=$SLACK_API $DEPLOY $@"
echo $CMD

eval $CMD
