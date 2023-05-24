# STACKIT DNS CLI Client
Community Client to interact with the STACKIT DNS API.
For early adopters who want to use the Service without the Portal integration yet avaliable.

## Feature List
+ Create
    + Create new DNS Zones
    + Create Record Set
+ Get
    + List single DNS Zone by ID
    + List all DNS Zone
    + List all Records
+ Delete
    + Delete DNS Zones
    + Delete Record Set

## Options Overview
```console
A command line interface for interacting with the DNS API.

Usage:
  dns [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  configure   Creates the configuration file needed to interact with the DNS API.
  create      Creates, updates and deletes resources for DNS.
  delete      Deletes DNS resources.
  get         Returns resources from DNS Zones.
  help        Help about any command

Flags:
      --authentication-token string   The JWT token for authenticating with the DNS API.
      --dns-api-url string            The url to the DNS API. (default "https://dns.api.stackit.cloud")
  -h, --help                          help for dns
      --project-id string             The project UUID the DNS resources are contained.

Use "dns [command] --help" for more information about a command.
```