
# snsbridge
This is a simple HTTP service that will automatically complete the Subscribe flow to start an SNS subscription from any topic and will put the records it recieves into a Kinesis stream.

## Significant Environment Variables (Configuration)
| Environment Variable | Default Value | Description |
| ----------- | ----------- | ----------- |
| AWS_ACCESS_KEY_ID | None | Access Key ID (standard AWS SDK variables) |
| AWS_SECRET_ACCESS_KEY | None | Access Key Secret (standard AWS SDK variables) |
| AWS_REGION | us-east-1 | Region for our Kinesis Stream |
| STREAM | TestStream | Kinesis Stream Name to push records to |
| PORT | 8080 | Which port should our HTTP server listen on? |
