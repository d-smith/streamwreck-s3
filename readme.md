# streamwreck s3

Reconcile the records processed from a stream from those written to a stream.

## Overview

idea - record processed items in an s3 bucket

## Deploy, invoke, etc.


Build and deploy the lambdas:

```console
make
make deploy
```

Write records to yonder stream:

```console
aws kinesis put-record --stream-name WreakStreamS3-Dev --data Data1234Foobar --partition-key foo
```

Get logs

```console
sam logs -n StreamProcessor --stack-name streamwrecks3
```
