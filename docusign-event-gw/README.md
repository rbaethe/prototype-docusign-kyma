# Overview

## Prerequisites

* Requires `go` setup.

## Local Run

* Build docker image

  ```shell script
  make build-image
  ```
* Run

  ```shell script
  docker run -p 8080:8080 gabbi/docusign-event-gw:0.0.4 --basic-auth-enabled=false --app-name=docusign --event-publish-url=http://localhost:8011/post
    
  docker run -p 8080:8080 rbdock1407/kyma-docusign:0.0.3 --basic-auth-enabled=false --app-name=docusign --event-publish-url=http://host.docker.internal:8011/post
   
  ```

* Verify

  ```shell script
  curl -X POST http://localhost:8080/events -d '{"key" : "value"}'
  ```

## Kyma Cluster

* Create Secret required for basic auth

  ```shell script
  kc -n <ns> create secret generic docusign-event-gw --from-literal=USERNAME=<user> --from-literal=PASSWORD=<password>
  ```

* Deploy service and deployment

  ```shell script
  kubectl apply -f k8s/
  ```
* Create an API from the Service.
* Note down the `URL`.

## DocuSign Side setup

* Set up a `Custom Connect Configuration` in DocuSign Admin console.
  * Use the `URL+'/events'` in the configuration.
  * Provide BasicAuth If configured in the event gateway.
