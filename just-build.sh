#!/bin/bash

NAME=${1:-min-instance}
TAG=${2:-0.01}
gcloud alpha builds submit --pack=image=us-central1-docker.pkg.dev/shingo-cloud-run-alwayson/myrepo/$NAME:$TAG,env=NAME=$NAME