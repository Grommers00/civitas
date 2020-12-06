# Civitas Backend

## Requirements

- Go 1.15
- AWS CLI 2.X
- AWS SAM CLI
- Docker
- Make

## Local Development

- Run `aws configure` and provide your AWS secret which can be found under `IAM`.
- Run `sam build` (or if you're on a \*Nix environment, `make`) to build the functions defined in `template.yaml`
- Run `sam local start-api --template template.yaml` to start a local developer server. This can be accessed at it's default `localhost:3000`.

**NOTE** Due to the current architecture of utilizing Lambdas, whereas this would be best as a REST API, we have to utilize a single endpoint per function handler.
So, instead of `/profile/{ID}`, we're passing in a url param (ex. `/profile?id={ID}`) directly to the endpoint. This is prevelant on all endpoints.

## References

- https://medium.com/better-programming/how-to-deploy-a-local-serverless-application-with-aws-sam-b7b314c3048c
