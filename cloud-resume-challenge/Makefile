.PHONY: build

build:
	sam build

deploy-infra:
	sam build --use-container && aws-vault exec my-user --no-session -- sam deploy

deploy-site:
	aws-vault exec my-user --no-session -- aws s3 sync ./resume-site s3://my-fantastic-website8417

invoke-put:
	sam build --use-container && aws-vault exec 