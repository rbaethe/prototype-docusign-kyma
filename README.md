# Overview

A sample integration of [DocuSign](https://www.docusign.com/) with [Kyma](https://kyma-project.io) to demonstrate how Kyma can be leveraged to enable side-by-side extensibility flows for DocuSign when used with SAP applications such as SAP Service Cloud.

It uses a [DocuSign addons connector](./addons) to connect DocuSign with Kyma.

## Demo workshop

### What will be covered?
* Set up connectivity between DocuSign and Kyma
* Deploy an example lambda that
  * Is executed on DocuSign Envelope events
  * Makes API calls to DocSign

### Steps

* 
![](./assets/steps/add-new-ns.png)
*
![](./assets/steps/create-ns.png)
*
![](./assets/steps/to-main.png)
![](./assets/steps/create-application.png)
![](./assets/steps/bind-application-to-ns.png)
*
![](./assets/steps/back-to-ns.png)
![](./assets/steps/addons-config.png)
![](./assets/steps/catalog-docusign-connector.png)
![](./assets/steps/provision-connector.png)
![](./assets/steps/apis-and-events.png)
![](./assets/steps/events-add-once.png)
apis add once
![](./assets/steps/docusign-add-connect-config.png)
![](./assets/steps/docusign-create-connect.png)
![](./assets/steps/docusign-connect-events.png)
create lambda
![](./assets/steps/lambda-event-trigger.png)
![](./assets/steps/lambda-code-dependencies.png)
![](./assets/steps/do-service-binding.png)
save lambda
create an envelope
check logs