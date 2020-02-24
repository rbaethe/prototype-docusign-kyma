---
title: Overview
type: Details
---

Use **DocuSign Connector** to connect DocuSign Events with Kyma events to trigger serverless compute and make API calls from a lambda or a microservice to DocuSign.

## Prerequisite

It uses the API token for making API calls to DocuSign.

## Components

![architecture](./assets/docusign-connector.png)

### API Registration Job

* Registers the [DocuSign APIs](https://raw.githubusercontent.com/docusign/eSign-OpenAPI-Specification/master/esignature.rest.swagger-v2.json) in Kyma application.
* Developers can then make API calls to DocuSign from microservcies and lambdas in Kyma.

### DocuSign Events Gateway

* Exposes an HTTP endpoint (with Basic Auth) which is required to registered as a **URL to Publish** in DocuSign Connect configuration.
* The gateway receives DocuSign Events, converts it to Kyma Events and forwards to the Kyma Event Service.
* The events can be configured in DocuSign to trigger microservices and lambdas in Kyma.
* It will be exposed as a Kyma API `https://docusign-event-gw-<Kyma namespace>-<kyma-cluster-domain>`. It will be visible under `Configuration -> APIs`