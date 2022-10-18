# cesiumts-lambda

Deploy [Cesium Terrain Server](https://github.com/geo-data/cesium-terrain-server) on AWS Lambda.

Uses [AWS Lambda Go API Proxy](https://github.com/awslabs/aws-lambda-go-api-proxy).

## Deploy

You will need to have installed AWS SAM CLI.

First build the image:

```
sam build
```

Then, deploy:

```
sam deploy --guided
```

Now follow the instructions.  In the end, SAM will have created all the
resources it needs.

## Contributing

Bug reports and pull requests are welcome on GitHub at the [issues
page](https://github.com/dymaxionlabs/cesiumts-lambda). This
project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [Contributor
Covenant](http://contributor-covenant.org) code of conduct.

## License

This project is licensed under Apache 2.0. Refer to
[LICENSE](https://github.com/dymaxionlabs/cesiumts-lambda/blob/main/LICENSE).
