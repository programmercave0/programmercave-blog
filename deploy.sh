#!/bin/bash

# Navigate to the Hugo site directory
hugo -t cayman-hugo-theme

# Deploy the public folder
cd public
git add .
git commit -m "new changes"
git push

# Return to the main directory and push the main repo
cd ..
git add .
git commit -m "new changes"
git push
