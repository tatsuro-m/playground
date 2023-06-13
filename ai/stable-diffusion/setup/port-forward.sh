#!/bin/bash

gcloud compute ssh stable-diffusion-main --zone=asia-northeast1-a -- -L 7860:localhost:7860
