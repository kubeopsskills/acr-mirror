# Azure Container Registry Mirror Tool

This project will help you to mirror Azure Container Registry across subscriptions.

It's possible to filter by docker tags.

<!-- TOC -->

- [acr-mirror](#acr-mirror)
    - [Install / Building](#install-building)
    - [Using](#using)
        - [Update all repositories](#update-all-repositories)
    - [Example config.yaml](#example-configyaml)

<!-- /TOC -->

## Install / Building

- make sure you got Go 1.16 or newer
  - OSX: `brew install go`
  - Windows: `winget install -e --id Golang.Go-Unstable`
- make sure you have `CGO` enabled
  - `export CGO_ENABLED=1`
- clone this repository to `$HOME/src/github.com/kubeopsskills/acr-mirror/cmd/acr`
- change your working directory to `$HOME/src/github.com/kubeopsskills/acr-mirror/cmd/acr`
- run `go install` to build and install the `acr-mirror` binary into your `$HOME/go/bin/` directory
  - alternative: `go build` to build the binary and put it in the current working directory

## Using

Make sure that you are logged into to `Azure` (`az login`)

`acr-mirror` will automatically mirror Azure Container Registry across subscriptions, so you do not need to login and do any UI operations in the Azure Portal.

### Update all repositories

- run `acr-mirror` and wait (for a while)

## Example config.yaml

```yml
---
source_subscription_id: "65113025-94f4-47c9-907a-81f3659078e6" # Azure source subscription id
target_subscription_id: "2c2288a3-11a6-4ce6-ae3a-28a336a1e86b" # Azure target subscription id

registries:
- name: "acr1oam"
  resourceGroupName: "TestACR1"
  repositories: 
  - "nginx"
  tags: # specific tag match
  - "*"
  target: 
  - name: "acr2oam"
    resourceGroupName: "TestACR2"
    prefix: "pregolden" # prefix for target repository
    rePrefix: "golden" # rename prefix for target repository
```

## Environment Variables

Environment Variable  |  Default       | Description
----------------------| ---------------| -------------------------------------------------
CONFIG_FILE_NAME      | config.yaml    | config file to use
RESOURCE_URL          | https://management.azure.com | Azure resource url