<p align="center">
<a href="https://www.tmone.com.my/"><img src="https://www.tmone.com.my/wp-content/themes/TM_One_Theme/images/tmone-logo.svg" width="82px" height="56px"></a>
</p>

<h1 align="center">TM Cloud Go Software Development Kit (Go SDK)</h1>

The TM Cloud Go SDK allows you to easily work with TM Cloud services such as Elastic Compute Service (ECS) and
Virtual Private Cloud (VPC) without the need to handle API related tasks.

This document introduces how to obtain and use TM Cloud Go SDK.

## Requirements

- To use TM Cloud Go SDK, you must have TM Cloud account as well as the Access Key (AK) and Secret key (SK) of the TM
  Cloud account. You can create an AccessKey in the TM Cloud console.

- To use TM Cloud Go SDK to access the APIs of specific service, please make sure you do have activated the service
  in [TM Cloud console](https://console.alphaedge.tmone.com.my/) if needed.

- TM Cloud Go SDK requires go 1.14 or later, run command `go version` to check the version of Go.

## Install Go SDK

Run the following command to install TM Cloud Go SDK:

``` bash
# Install the library of TM Cloud Go SDK
go get github.com/tmcloud-sdk/tmcloud-sdk-go
```

## Code Example

- The following example shows how to query a list of VPCs in a specific region, you need to substitute your
  real `{service} "github.com/tmcloud-sdk/tmcloud-sdk-go/services/{service}/{version}"`
  for `vpc "github.com/tmcloud-sdk/tmcloud-sdk-go/services/vpc/v2"` in actual use, and initialize the client
  as `{service}.New{Service}Client`.
- Substitute the values for `{your ak string}`, `{your sk string}`, `{your endpoint string}` and `{your project id}`.

``` go
package main

import (
    "fmt"
    "github.com/tmcloud-sdk/tmcloud-sdk-go/core/auth/basic"
    "github.com/tmcloud-sdk/tmcloud-sdk-go/core/config"
    "github.com/tmcloud-sdk/tmcloud-sdk-go/core/httphandler"
    vpc "github.com/tmcloud-sdk/tmcloud-sdk-go/services/vpc/v2"
    "github.com/tmcloud-sdk/tmcloud-sdk-go/services/vpc/v2/model"
    "net/http"
)

func RequestHandler(request http.Request) {
    fmt.Println(request)
}

func ResponseHandler(response http.Response) {
    fmt.Println(response)
}

func main() {
    client := vpc.NewVpcClient(
        vpc.VpcClientBuilder().
            WithEndpoint("{your endpoint}").
            WithCredential(
                basic.NewCredentialsBuilder().
                    WithAk("{your ak string}").
                    WithSk("{your sk string}").
                    WithProjectId("{your project id}").
                    Build()).
            WithHttpConfig(config.DefaultHttpConfig().
                WithIgnoreSSLVerification(true).
                WithHttpHandler(httphandler.
                    NewHttpHandler().
                        AddRequestHandler(RequestHandler).
                        AddResponseHandler(ResponseHandler))).
            Build())

    limit := int32(1)
    request := &model.ListVpcsRequest{
      Limit: &limit,
    }
    response, err := client.ListVpcs(request)
    if err == nil {
      fmt.Printf("%+v\n\n", response.Vpcs)
    } else {
      fmt.Println(err)
    }
}
```

## Changelog

Detailed changes for each released version are documented in
the [CHANGELOG.md](https://github.com/tmcloud-sdk/tmcloud-sdk-go/blob/master/CHANGELOG.md).

## User Manual [:top:](#tm-cloud-go-software-development-kit-go-sdk)

* [1. Client Configuration](#1-client-configuration-top)
    * [1.1 Default Configuration](#11-default-configuration-top)
    * [1.2 Network Proxy](#12-network-proxy-top)
    * [1.3 Timeout Configuration](#13-timeout-configuration-top)
    * [1.4 SSL Certification](#14-ssl-certification-top)
    * [1.5 Custom Network Connection](#15-custom-network-connection-top)
* [2. Credentials Configuration](#2-credentials-configuration-top)
  * [2.1 Use Permanent AK&SK](#21-use-permanent-aksk-top)
  * [2.2 Use Temporary AK&SK](#22-use-temporary-aksk-top)
* [3. Client Initialization](#3-client-initialization-top)
  * [3.1 Initialize client with specified Endpoint](#31-initialize-the-serviceclient-with-specified-endpoint-top)
* [4. Send Request and Handle response](#4-send-requests-and-handle-responses-top)
    * [4.1 Exceptions](#41-exceptions-top)
* [5.Troubleshooting](#5-troubleshooting-top)
    * [5.1 Original HTTP Listener](#51-original-http-listener-top)
* [6. Upload and download files](#6-upload-and-download-files-top)
* [7. Retry For Request](#7-retry-for-request-top)

### 1. Client Configuration [:top:](#user-manual-top)

#### 1.1 Default Configuration [:top:](#user-manual-top)

``` go
import "github.com/tmcloud-sdk/tmcloud-sdk-go/core/config"

// Use default configuration
httpConfig := config.DefaultHttpConfig()
```

#### 1.2 Network Proxy [:top:](#user-manual-top)

``` go
// Use proxy if needed
httpConfig.WithProxy(config.NewProxy().
    WithSchema("http").
    WithHost("proxy.tmone.com.my").
    WithPort(80).
    WithUsername("testuser").
    WithPassword("password"))))
```

#### 1.3 Timeout Configuration [:top:](#user-manual-top)

``` go
// The default timeout is 120 seconds, which can be adjusted as needed
httpConfig.WithTimeout(120);
```

#### 1.4 SSL Certification [:top:](#user-manual-top)

``` go
// Skip SSL certification checking while using https protocol if needed
httpConfig.WithIgnoreSSLVerification(true);
```

#### 1.5 Custom Network Connection [:top:](#user-manual-top)

``` go
// Config network connection dial function if needed
func DialContext(ctx context.Context, network string, addr string) (net.Conn, error) {
	return net.Dial(network, addr)
}
httpConfig.WithDialContext(DialContext)
```

### 2. Credentials Configuration [:top:](#user-manual-top)

There are two types of TM Cloud services, `regional` services and `global` services.

Global services contain IAM.

For `regional` services' authentication, projectId is required to initialize basic.NewCredentialsBuilder(). 

For `global` services' authentication, domainId is required to initialize global.NewCredentialsBuilder().

The following authentications are supported:

- permanent AK&SK
- temporary AK&SK + SecurityToken

#### 2.1 Use Permanent AK&SK [:top:](#user-manual-top)

**Parameter description**:

- `ak` is the access key ID for your account.
- `sk` is the secret access key for your account.
- `projectId` is the ID of your project depending on your region which you want to operate.
- `domainId` is the account ID of TM Cloud.

``` go
// Regional Services
basicCredentials := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    Build()

// Global Services
globalCredentials := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithDomainId(domainId).
    Build()
```

#### 2.2 Use Temporary AK&SK [:top:](#user-manual-top)

It's required to obtain temporary AK&SK and security token first, which could be obtained through
permanent AK&SK or through an agency.A temporary access key and securityToken are issued by the system to IAM users, and can be valid for 15 minutes to 24 hours.

- Obtaining a temporary access key and security token through token, you could refer to
document: [Obtaining a Temporary AK/SK](https://support.alphaedge.tmone.com.my/en-us/api/iam/en-us_topic_0097949518.html). The API mentioned in the document above
corresponds to the method of `CreateTemporaryAccessKeyByToken` in IAM SDK.

**Parameter description**:

- `ak` is the access key ID for your account.
- `sk` is the secret access key for your account.
- `securityToken` is the security token when using temporary AK/SK.
- `projectId` is the ID of your project depending on your region which you want to operate.
- `domainId` is the account ID of TM Cloud.

``` go
// Regional Services
basicCredentials := basic.NewCredentialsBuilder().
            WithAk(ak).
            WithSk(sk).
            WithProjectId(projectId).
            WithSecurityToken(securityToken).
            Build()

// Global Services
globalCredentials := global.NewCredentialsBuilder().
            WithAk(ak).
            WithSk(sk).
            WithDomainId(domainId).
            WithSecurityToken(securityToken).
            Build()
```

### 3. Client Initialization [:top:](#user-manual-top)

#### 3.1 Initialize the {Service}Client with specified Endpoint [:top:](#user-manual-top)

``` go
// Specify the endpoint, take the endpoint of VPC service in region of cn-north-4 for example
endpoint := "${endpoint}"

// Initialize the credentials, you should provide projectId or domainId in this way, take initializing BasicCredentials for example
basicAuth := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    Build()

// Initialize specified New{Service}Client, take initializing the regional service VPC's VpcClient for example
client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
        WithEndpoint(endpoint).
        WithCredential(basicCredentials).
        WithHttpConfig(config.DefaultHttpConfig()).  
        Build())
```

**where:**

- `endpoint` varies with services and regions,
  see [Regions and Endpoints](https://support.alphaedge.tmone.com.my/zh-cn/endpoint/index.html) to obtain correct endpoint.

### 4. Send Requests and Handle Responses [:top:](#user-manual-top)

``` go
// send a request and print response, take interface of ListVpcs for example
limit := int32(1)
request := &model.ListVpcsRequest{
    Limit: &limit,
}

response, err := client.ListVpcs(request)
if err == nil {
    fmt.Printf("%+v\n\n", response.Vpcs)
} else {
    fmt.Println(err)
}
```

#### 4.1 Exceptions [:top:](#user-manual-top)

| Level 1 | Notice |
| :---- | :---- |
| ServiceResponseError | service response error |
| url.Error | connect endpoint error |

``` go
response, err := client.ListVpcs(request)
if err == nil {
    fmt.Printf("%+v\n\n", response.Vpcs)
} else {
    fmt.Println(err)
}
```

### 5. Troubleshooting [:top:](#user-manual-top)

#### 5.1 Original HTTP Listener [:top:](#user-manual-top)

In some situation, you may need to debug your http requests, original http request and response information will be
needed. The SDK provides a listener function to obtain the original encrypted http request and response information.

> :warning:  Warning: The original http log information is used in debugging stage only, please do not print the original http header or body in the production environment. This log information is not encrypted and contains sensitive data such as the password of your ECS virtual machine, or the password of your IAM user account, etc. When the response body is binary content, the body will be printed as "***" without detailed information.

``` go
func RequestHandler(request http.Request) {
    fmt.Println(request)
}

func ResponseHandler(response http.Response) {
    fmt.Println(response)
}

client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
        WithEndpoint("{your endpoint}").
        WithCredential(
            basic.NewCredentialsBuilder().
                WithAk("{your ak string}").
                WithSk("{your sk string}").
                WithProjectId("{your project id}").
                   Build()).
        WithHttpConfig(config.DefaultHttpConfig().
            WithIgnoreSSLVerification(true).
            WithHttpHandler(httphandler.
                NewHttpHandler().
                    AddRequestHandler(RequestHandler).
                    AddResponseHandler(ResponseHandler))).
        Build())
```

### 6. Upload and download files [:top:](#user-manual-top)

```go
package main

import (
	"fmt"
	"os"

	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/auth/basic"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/def"

	service "github.com/tmcloud-sdk/tmcloud-sdk-go/services/service/v1"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/service/v1/model"
)

func uploadAndDownloadFile(client *service.ServiceClient) {
	
	// Open the file.
	file, err := os.Open("demo.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &model.UploadFileRequestBody{
		File:           def.NewFilePart(file),
		FileName: def.NewMultiPart("rename.jpg"),
	}

	request := &model.UploadFileRequest{Body: body}
	response, err := client.UploadFile(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
		return
	}

	// Download the file.
	result, err := os.Create("result.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	response.Consume(result)

}

func main() {
	ak := "{your ak string}"
	sk := "{your sk string}"
	endpoint := "{your endpoint string}"
	projectId := "{your project id}"

	credentials := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		WithProjectId(projectId).
		Build()

	client := service.NewServiceClient(
		service.ServiceClientBuilder().
			WithEndpoint(endpoint).
			WithCredential(credentials).
			Build())

    uploadAndDownloadFile(client)
}
```

### 7. Retry For Request [:top:](#user-manual-top)

When a request encounters a network exception or flow control on the interface, the request needs to be retried. The
Go SDK provides the retry method for our users which could be used to the requests of `GET` HTTP method. 
If you want to use the retry method, the following parameters are required:

- _maxRetryTimes_: the max retry times
- _retryCondition_: a function, which determine the condition of when to retry
- _backoffStrategy_: calculate the wait duration before next retry

Take the interface `ListVpcs` of VPC service for example, assume the request would retry at most 3 times, 
retry when service responses an error, the code would be like the following:

``` go
// initialize the client
client := vpc.NewVpcClient(
	vpc.VpcClientBuilder().
		WithEndpoint("<input your endpoint>").
		WithCredential(
			basic.NewCredentialsBuilder().
				WithAk("<input your ak>").
				WithSk("<input your sk>").
				WithProjectId("<input your project id>").
				Build()).
		Build())

// initialize the request
request := &model.ListVpcsRequest{}

// send the requet and retry when service responses an error
response, err := client.ListVpcsInvoker(request).WithRetry(3, func(i interface{}, err error) bool {
	return err != nil
}, new(retry.None)).Invoke()

if err == nil {
	fmt.Printf("%+v\n", response)
} else {
	fmt.Printf("%+v\n", err)
}
```
