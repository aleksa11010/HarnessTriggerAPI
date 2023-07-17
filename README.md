# Harness Trigger API 

This project provides a Go program that uses the Harness API to trigger workflows and pipelines based on specified YAML files.

## Prerequisites 

- Go 1.16 or higher
- An account with Harness

## Dependencies

- `github.com/go-resty/resty/v2`: A simple HTTP and REST client library for Go.
- `github.com/sirupsen/logrus`: A structured logger for Go.

## How to Run

The program takes two flags, `-config` and `-trigger`, which specify the paths to the Account Config and Trigger YAML files, respectively.

```
go run main.go -config <config-file-path> -trigger <trigger-file-path>
```

Replace `<config-file-path>` and `<trigger-file-path>` with the paths to your YAML files.

If you don't provide these flags when running the program, it will exit and log an error message.

The `config` YAML file should provide the following details:

- `accountIdentifier`: Account Identifier for your Harness account.
- `orgIdentifier`: Organization Identifier for your Harness account.
- `projectIdentifier`: Project Identifier for your Harness account.
- `targetIdentifier`: Target Identifier for your Harness account.
- `apiKey` : API Key for your Harness account.

The `trigger` YAML file should contain the YAML configuration of the trigger you want to create. It can contain placeholder values for some fields in the format `<+PLACEHOLDER>`. These placeholders will be replaced with real values from your `config` file. The following placeholders are supported:

- `<+ORG>`: Will be replaced with the `orgIdentifier` from your `config` file.
- `<+PROJECT>`: Will be replaced with the `projectIdentifier` from your `config` file.
- `<+PIPELINE>`: Will be replaced with the `targetIdentifier` from your `config` file.

The program will send a POST request to the `/pipeline/api/triggers` endpoint of the Harness API, using the API key and details from the `config` YAML file, and the trigger details from the `trigger` YAML file. 

If there's an error during the API request, the program will log the error. It will also log the response from the API request.

## Examples

In the `examples` directory, you can find examples of `config` and `trigger` YAML files that can be used with this program. You can use these as a starting point for creating your own configuration files.

Here is how you can run the program with these example files:

```
go run main.go -config examples/config.yaml -trigger examples/trigger.yaml
```

Replace `examples/config.yaml` and `examples/trigger.yaml` with the paths to your own YAML files if you have created custom configurations.
