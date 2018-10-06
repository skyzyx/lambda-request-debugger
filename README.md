# Request Debugger for AWS Lambda

Will dump the request data that AWS Lambda receives. This is very similar to services like [RequestBin](https://requestbin.fullcontact.com), but will show you the data that you receive from API Gateway and the AWS Lambda systems.

This is a _Serverless_ app, written in Go ([Golang]), running in AWS Lambda, with API Gateway in front of it, and AWS CloudFront in front of that (for caching).

**This is an experiment.** Uptime is not guaranteed, and there is no SLA. _But_ all-in-all, it should be reasonably reliable.

> **NOTE:** I've tried to build this in a way that is very cheap to run. But if you find yourself using this more than just occasionally, consider kicking me down a few bucks to offset the cost of running this service. <https://cash.me/$rparman>

## Usage

The `https://debug.ryanparman.com` hostname is a CloudFront distribution, in front of API Gateway, in front of a Lambda function.

There are two endpoints:

* `/json` — Send an HTTP request to this endpoint, and the response will be in JSON format.
* `/dump` — Send an HTTP request to this endpoint, and the response will be in a custom _variable dump_ format from [go-spew](https://github.com/davecgh/go-spew#sample-dump-output). This allows you to see the underlying Golang types.

It responds to `GET`, `POST`, `PUT`, `PATCH`, and `DELETE` HTTP verbs.

## Examples

### JSON

#### Request

```bash
curl -X "POST" "https://debug.ryanparman.com/json" \
     -H 'Content-Type: text/plain; charset=utf-8' \
     -d "I am posting a request body."
```

#### Response

```http
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
X-Amz-Cf-Id: RhfUXoXBJEcjWHNSC1B09BSR8AT48pYrtGAeXNXTeUgh6FF4xaojGA==
X-Amzn-Trace-Id: Root=1-5bb818a8-1425e9aaa604a1e7de1f8a2c;Sampled=0
X-Cache: Miss from cloudfront
Via: 1.1 b2532cb29a55e8fe8106a4a9a9241592.cloudfront.net (CloudFront), 1.1 19f9923c4e449b92312c8813bf9135f5.cloudfront.net (CloudFront)
Expires: Sat, 06 Oct 2018 02:06:33 GMT
x-amz-apigw-id: OUjKYEjYoAMFYWg=
Date: Sat, 06 Oct 2018 02:06:33 GMT
Content-Length: 1929
Connection: keep-alive
x-amzn-RequestId: 7243dabe-c90c-11e8-81ac-333fe7e5da01
Last-Modified: Sat, 06 Oct 2018 02:06:33 GMT

{
    "resource": "/json",
    "path": "/json",
    "httpMethod": "POST",
    "headers": {
        "Accept-Encoding": "br, gzip, deflate",
        "CloudFront-Forwarded-Proto": "https",
        "CloudFront-Is-Desktop-Viewer": "true",
        "CloudFront-Is-Mobile-Viewer": "false",
        "CloudFront-Is-SmartTV-Viewer": "false",
        "CloudFront-Is-Tablet-Viewer": "false",
        "CloudFront-Viewer-Country": "US",
        "Content-Type": "application/json; charset=utf-8",
        "Host": "hq8su3sz27.execute-api.us-east-1.amazonaws.com",
        "User-Agent": "Amazon CloudFront",
        "Via": "1.1 19f9923c4e449b92312c8813bf9135f5.cloudfront.net (CloudFront), 1.1 b2532cb29a55e8fe8106a4a9a9241592.cloudfront.net (CloudFront)",
        "X-Amz-Cf-Id": "a4HZKxs9jPW9dFuQkXrhYFMArl7wSxBjj3TQG06YKKKvln7CuIHCBQ==",
        "X-Amzn-Trace-Id": "Root=1-5bb818a8-1425e9aaa604a1e7de1f8a2c",
        "X-Forwarded-For": "2601:601:957f:e6c2:cdba:b9db:8e2a:3dff, 52.46.16.95, 52.46.16.65",
        "X-Forwarded-Port": "443",
        "X-Forwarded-Proto": "https"
    },
    "queryStringParameters": null,
    "pathParameters": null,
    "stageVariables": null,
    "requestContext": {
        "accountId": "133904017518",
        "resourceId": "d3gfbs",
        "stage": "dev",
        "requestId": "7243dabe-c90c-11e8-81ac-333fe7e5da01",
        "identity": {
            "cognitoIdentityPoolId": "",
            "accountId": "",
            "cognitoIdentityId": "",
            "caller": "",
            "apiKey": "",
            "sourceIp": "52.46.16.95",
            "cognitoAuthenticationType": "",
            "cognitoAuthenticationProvider": "",
            "userArn": "",
            "userAgent": "Amazon CloudFront",
            "user": ""
        },
        "resourcePath": "/json",
        "authorizer": null,
        "httpMethod": "POST",
        "apiId": "hq8su3sz27"
    },
    "body": "I am posting a request body."
}
```

### Dump

#### Request

```bash
curl -X "POST" "https://debug.ryanparman.com/dump" \
     -H 'Content-Type: text/plain; charset=utf-8' \
     -d "I am posting a request body."
```

#### Response

```http
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
X-Amz-Cf-Id: CR0uIbaiVn9x0VPgr7K5CvX4C4-CF_hsOGcfO2b7i3abOIWlu76wVg==
X-Amzn-Trace-Id: Root=1-5bb81a27-4270b1c8e9c91744a4dbcdd1;Sampled=0
X-Cache: Miss from cloudfront
Via: 1.1 33cfbeb7154bbef1432b207659c6dac5.cloudfront.net (CloudFront), 1.1 2d845616bb025756fb00e75135987ac7.cloudfront.net (CloudFront)
Expires: Sat, 06 Oct 2018 02:12:55 GMT
x-amz-apigw-id: OUkGOGfDoAMFcIg=
Date: Sat, 06 Oct 2018 02:12:55 GMT
Content-Length: 3001
Connection: keep-alive
x-amzn-RequestId: 568cb656-c90d-11e8-940e-3352e72e7d24
Last-Modified: Sat, 06 Oct 2018 02:12:55 GMT

(events.APIGatewayProxyRequest) {
    Resource: (string) (len=5) "/dump",
    Path: (string) (len=5) "/dump",
    HTTPMethod: (string) (len=4) "POST",
    Headers: (map[string]string) (len=16) {
        (string) (len=15) "Accept-Encoding": (string) (len=17) "br, gzip, deflate",
        (string) (len=12) "Content-Type": (string) (len=25) "text/plain; charset=utf-8",
        (string) (len=4) "Host": (string) (len=46) "hq8su3sz27.execute-api.us-east-1.amazonaws.com",
        (string) (len=16) "X-Forwarded-Port": (string) (len=3) "443",
        (string) (len=27) "CloudFront-Is-Mobile-Viewer": (string) (len=5) "false",
        (string) (len=28) "CloudFront-Is-SmartTV-Viewer": (string) (len=5) "false",
        (string) (len=25) "CloudFront-Viewer-Country": (string) (len=2) "US",
        (string) (len=3) "Via": (string) (len=130) "1.1 2d845616bb025756fb00e75135987ac7.cloudfront.net (CloudFront), 1.1 33cfbeb7154bbef1432b207659c6dac5.cloudfront.net (CloudFront)",
        (string) (len=11) "X-Amz-Cf-Id": (string) (len=56) "pmeK0suyD_AZp_5zWUhGCEN5T3xU8f24bDbiosYrDrjoKUlpA7MCRg==",
        (string) (len=26) "CloudFront-Forwarded-Proto": (string) (len=5) "https",
        (string) (len=28) "CloudFront-Is-Desktop-Viewer": (string) (len=4) "true",
        (string) (len=15) "X-Amzn-Trace-Id": (string) (len=40) "Root=1-5bb81a27-4270b1c8e9c91744a4dbcdd1",
        (string) (len=15) "X-Forwarded-For": (string) (len=64) "2601:601:957f:e6c2:cdba:b9db:8e2a:3dff, 52.46.16.95, 52.46.16.95",
        (string) (len=17) "X-Forwarded-Proto": (string) (len=5) "https",
        (string) (len=27) "CloudFront-Is-Tablet-Viewer": (string) (len=5) "false",
        (string) (len=10) "User-Agent": (string) (len=17) "Amazon CloudFront"
    },
    QueryStringParameters: (map[string]string) <nil>,
    PathParameters: (map[string]string) <nil>,
    StageVariables: (map[string]string) <nil>,
    RequestContext: (events.APIGatewayProxyRequestContext) {
        AccountID: (string) (len=12) "133904017518",
        ResourceID: (string) (len=6) "bgf9uc",
        Stage: (string) (len=3) "dev",
        RequestID: (string) (len=36) "568cb656-c90d-11e8-940e-3352e72e7d24",
        Identity: (events.APIGatewayRequestIdentity) {
            CognitoIdentityPoolID: (string) "",
            AccountID: (string) "",
            CognitoIdentityID: (string) "",
            Caller: (string) "",
            APIKey: (string) "",
            SourceIP: (string) (len=11) "52.46.16.95",
            CognitoAuthenticationType: (string) "",
            CognitoAuthenticationProvider: (string) "",
            UserArn: (string) "",
            UserAgent: (string) (len=17) "Amazon CloudFront",
            User: (string) ""
        },
        ResourcePath: (string) (len=5) "/dump",
        Authorizer: (map[string]interface {}) <nil>,
        HTTPMethod: (string) (len=4) "POST",
        APIID: (string) (len=10) "hq8su3sz27"
    },
    Body: (string) (len=28) "I am posting a request body.",
    IsBase64Encoded: (bool) false
}
```

## Developing/Deploying

### Golang

Go (when spoken) or [Golang] (when written) is a strongly-typed language from Google that "blends the simplicity of Python with the performance of C". Static binaries can be compiled for all major platforms, and many minor ones.

It is recommended that you install Golang using your system's package manager. If you don't have one (or if the version is too old), you can [install Golang from its website](https://golang.org/doc/install). Reading the [Getting Started](https://golang.org/doc/) documentation is a valuable exercise.

```bash
brew update && brew install golang
```

### Glide

Golang dependencies are managed with [Glide]. You should install them before compiling this project.

```bash
curl https://glide.sh/get | sh
glide install
```

### GoMetaLinter

[GoMetaLinter] pulls together many popular linting tools, and can run them on a project.

```bash
gometalinter.v2 --install
```

### Serverless

[Serverless] is a platform that wraps AWS Lambda and AWS CloudFormation, simplifying the deployment of Lambda apps. Serverless is written in Node.js, so you need to install that first.

I recommend you install the [Node Version Manager][nvm], and use that to install the latest Node.js and npm. Once that's complete, install `serverless`.

```bash
npm i -g serverless
```

### Developing

This app is small, and is self-contained in `main.go`.

_By default_, it expects to be running in AWS Lambda, receiving HTTP requests coming from API Gateway.

If you are performing local development/testing, run `make build` to build for the local platform.

```bash
make build
```

Make sure that you run the linter to catch any issues.

```bash
make lint
```

### Deployment

`serverless` uses the same [local credentials](https://docs.aws.amazon.com/cli/latest/topic/config-vars.html) that the AWS CLI tools and the AWS SDKs use. If you haven't configured those yet, do that first.

Run `make package` to build a binary for AWS Lambda. Then, `serverless deploy` to deploy the app to your environment.

```bash
make package
serverless deploy
```

  [Glide]: https://glide.sh
  [Golang]: https://golang.org
  [GoMetaLinter]: https://github.com/alecthomas/gometalinter
  [nvm]: https://github.com/creationix/nvm
  [Serverless]: https://serverless.com/framework/docs/getting-started/
