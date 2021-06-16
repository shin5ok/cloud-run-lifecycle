#!/bin/bash

gcloud beta run deploy --source=. --platform=managed --region=us-central1 --allow-unauthenticated cloud-run-lifecycle $@