#!/bin/bash

# Build site
hugo -t cayman-hugo-theme

# Deploy submodule
cd public
git checkout main       # switch to branch before commit
git pull                # update branch
git add .
git commit -m "new changes"
git push
cd ..

# Push main repo
git add .
git commit -m "new changes"
git push
