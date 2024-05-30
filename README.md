# Reverse proxy for AWS S3 w/ basic authentication

## Description

This is a reverse proxy for AWS S3, which is able to provide basic authentication as well.  
You don't need to configure a Bucket for `Website Hosting`.

http://this-proxy.com/access/ -> s3://bucket/access/index.html

If auth is enable, you can upload to the bucket (assuming your AWS credentials permit it). Strongly recommended to have SSL enabled for this, as the basic auth will be sent in plain text otherwise

```bash
 curl http://[::1]:21080/disco.gif -u${AUTH} --data-binary @disco.gif
 ```

## Usage

### Set environment variables

Environment Variables     | Description                                       | Required | Default
------------------------- | ------------------------------------------------- | -------- | -----------------
PRIMARY_STORE_ACCESS_KEY         | Primary AWS `access key` for API access.                  |          | EC2 Instance Role
PRIMARY_STORE_SECRET_KEY     | Primary AWS `secret key` for API access.                  |          | EC2 Instance Role
SECONDARY_STORE_ACCESS_KEY         | Secondary AWS `access key` for API access.                  |          | EC2 Instance Role
SECONDARY_STORE_SECRET_KEY     | Secondary AWS `secret key` for API access.                  |          | EC2 Instance Role

Other environment variables can be set by `S3_PROXY_` and uppercase CLI options without hyphens or underscores, so `--listen-port` becomes `S3_PROXY_LISTENPORT`.

### Set CLI options

```bash
$ aws-s3-proxy serve -h
serve the s3 proxy

Usage:
  aws-s3-proxy serve [flags]

Flags:
      --facility string                               Location where the service is running
      --healthcheck-path string                       path for healthcheck
  -h, --help                                          help for serve
      --http-cache-control Cache-Control              overrides S3's HTTP Cache-Control header
      --http-expires Expires                          overrides S3's HTTP Expires header
      --listen-address string                         host address to listen on (default "::1")
      --listen-port string                            port to listen on (default "21080")
      --primary-store-bucket string                   bucket name
      --primary-store-disable-bucket-ssl              toggle tls for the aws-sdk
      --primary-store-disable-compression             toggle compressions
      --primary-store-endpoint string                 endpoint URL (hostname only or fully qualified URI)
      --primary-store-idle-connection-timeout int     idle connection timeout in seconds (default 10)
      --primary-store-insecure-tls                    toogle tls verify
      --primary-store-max-idle-connections int        max idle connections (default 150)
      --primary-store-region string                   region for bucket
      --secondary-fall-back                           toggle read from secondary
      --secondary-store-bucket string                 bucket name
      --secondary-store-disable-bucket-ssl            toggle tls for the aws-sdk
      --secondary-store-disable-compression           toggle compressions
      --secondary-store-endpoint string               endpoint URL (hostname only or fully qualified URI)
      --secondary-store-idle-connection-timeout int   idle connection timeout in seconds (default 10)
      --secondary-store-insecure-tls                  toogle tls verify
      --secondary-store-max-idle-connections int      max idle connections (default 150)
      --secondary-store-region string                 region for bucket

Global Flags:
      --config string   config file (default is $HOME/.s3-proxy.yaml)
      --debug           Enable debug logging
      --pretty          Enable pretty (human readable) logging output
```

## Copyright and license

Code released under the [MIT license](https://github.com/packethost/aws-s3-proxy/blob/master/LICENSE).
