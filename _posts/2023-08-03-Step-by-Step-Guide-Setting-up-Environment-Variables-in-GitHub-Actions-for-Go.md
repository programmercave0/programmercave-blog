---
layout: post
title: "Step-by-Step Guide: Setting up Environment Variables in GitHub Actions for Go"
description: "GitHub Actions is a powerful platform that allows developers to automate workflows and build, test, and deploy their applications with ease. One common use case is to securely pass environment variables, such as API keys or passwords, to a Go program during the workflow execution in GitHub Actions. In this blog, we'll walk through the process of setting up repository secrets, creating a GitHub Actions workflow, and using those secrets as environment variables in a Go program."
author: "Programmercave"
header-img: "/assets/github-action-env-var/github-action-env-var.png"
tags:  [Github-Actions, Go, DevOps]
date: 2023-08-03
---
* toc
{:toc}

## Introduction:

GitHub Actions is a powerful platform that allows developers to automate workflows and build, test, and deploy their applications with ease. One common use case is to securely pass environment variables, such as API keys or passwords, to a Go program during the workflow execution in GitHub Actions. In this blog, we'll walk through the process of setting up repository secrets, creating a GitHub Actions workflow, and using those secrets as environment variables in a Go program.

![Setting up Environment Variables in GitHub Actions for Go]({{ site.url }}/assets/github-action-env-var/github-action-env-var.png){:class="img-responsive"}

## Creating Secrets in GitHub Repository:

 - Open your GitHub repository on the web.
 - Navigate to "Settings" in the right sidebar.
 - Click on "Secrets and variables" in the left sidebar and then "Actions".
 ![Setting up Environment Variables in GitHub Actions for Go]({{ site.url }}/assets/github-action-env-var/github-action-env-var-1.png){:class="img-responsive"}
 - Click on "New repository secret."
 ![Setting up Environment Variables in GitHub Actions for Go]({{ site.url }}/assets/github-action-env-var/github-action-env-var-2.png){:class="img-responsive"}
 - Enter the name of the secret (e.g., API_KEY) in the "Name" field.
 - Add the value of the secret (e.g., your actual API key) in the "Value" field.
 - Click "Add secret" to save it.
 ![Setting up Environment Variables in GitHub Actions for Go]({{ site.url }}/assets/github-action-env-var/github-action-env-var-3.png){:class="img-responsive"}

Repeat the process for any other sensitive information you need to store, like passwords or access tokens.

## Creating the Go Program:

For this example, let's assume we have a simple Go program that requires two sensitive environment variables, `API_KEY` and `DB_PASSWORD`. Create a new file named `main.go` in the root of your repository with the following content:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	apiKey := os.Args[1]
	dbPassword := os.Args[2]

	fmt.Println("API_KEY:", apiKey)
	fmt.Println("DB_PASSWORD:", dbPassword)

	// Your main program logic here...
}
```

## Creating the GitHub Actions Workflow:

1. In the root of your repository, create a new directory named `.github` if it doesn't exist.
2. Inside the `.github` directory, create another directory named `workflows`.
3. Create a new file named `main.yml` inside the `workflows` directory with the following content:

{% raw %}

```yaml
name: Execute Go Program

on:
  push:
    branches:
      - main

jobs:
  execute_go_program:
    runs-on: ubuntu-latest

    steps:
      - name: Check out repository
        uses: actions/checkout@v3
        with:
          persist-credentials: false

      - name: Set up environment and run Go program
        env:
          API_KEY: ${{ secrets.API_KEY }}
          DB_PASSWORD: ${{ secrets.DB_PASSWORD }}
        run: go run main.go "${{ env.API_KEY }}" "${{ env.DB_PASSWORD }}"
```

## Explanation of the Workflow:

1. The workflow will run on every push event to the `main` branch (`on: push: branches: [ "main" ]`).
2. It defines a single job named `execute_go_program`.
3. The job runs on the latest version of the Ubuntu operating system (`runs-on: ubuntu-latest`).
4. The first step checks out the repository (`uses: actions/checkout@v3`).
5. The second step sets up the environment variables `API_KEY` and `DB_PASSWORD` using repository secrets (`secrets.API_KEY` and `secrets.DB_PASSWORD`).
6. The third step runs the Go program using `go run main.go` and passes the environment variables as command-line arguments (`"${{ env.API_KEY }}"` and `"${{ env.DB_PASSWORD }}"`).

{% endraw %}
## Conclusion:

By following the steps above, you can securely pass environment variables to your Go program in GitHub Actions. The use of repository secrets ensures that sensitive information is kept private and not exposed in plaintext in your workflows. With this setup, you can confidently automate your Go programs with GitHub Actions while maintaining a high level of security for your sensitive data.

## FAQs:

1. **Q:** Can I use other programming languages with GitHub Actions?
   - Yes, GitHub Actions supports multiple programming languages and environments.

2. **Q:** Is there a limit to the number of secrets I can store in my GitHub repository?
   - GitHub allows you to store up to 100 secrets per repository.

3. **Q:** Can I access the secrets programmatically from my Go program?
   - Yes, as shown in the example, you can access secrets using environment variables in your Go program.

4. **Q:** How can I ensure that only authorized users can access the repository secrets?
   - GitHub provides access controls and permissions to manage who can view and use repository secrets.

5. **Q:** Are repository secrets encrypted at rest?
   - Yes, GitHub encrypts and securely stores repository secrets to ensure their protection.
