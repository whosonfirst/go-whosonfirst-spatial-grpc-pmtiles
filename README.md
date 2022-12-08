# go-whosonfirst-spatial-grpc-pmtiles

Go package implementing the `go-whosonfirst-spatial-grpc` toolchain with support for Protomaps PMTiles databases.

## Documentation

Documentation is incomplete at this time.

## Tools

### server

```
$> make cli
go build -mod vendor -o bin/server cmd/server/main.go
```

```
$> ./bin/server \
	-spatial-database-uri 'pmtiles://?tiles=file:///usr/local/data&database=whosonfirst'
```

Or if you're PMTiles database is hosted in an S3 bucket:

```
$> ./bin/server \
	-spatial-database-uri 'pmtiles://?tiles=s3blob%3A%2F%2Ftiles%3Fregion%3Dus-east-1%26prefix%3Dpip%2F%26credentials%3Dsession&database=whosonfirst'
```

And then in [whosonfirst/go-whosonfirst-spatial-grpc](https://github.com/whosonfirst/go-whosonfirst-spatial-grpc):

```
$> go build -mod vendor -o bin/client cmd/client/main.go

$> ./bin/client \
	-latitude 37.615761 \
	-longitude -122.389154 \
	-sort-uri placetype:// \
	-sort-uri name:// \
	-sort-uri inception:// \
	
	| jq '.places[] | "\(."name") \(."inception_date") (\(."placetype"))"'
	
"North America null (continent)"
"United States of America null (empire)"
"United States null (country)"
"America/Los_Angeles null (timezone)"
"SFO Terminal Complex 2017~ (building)"
"SFO Terminal Complex 2021-05-25 (building)"
"SFO Terminal Complex 2021-11-09 (building)"
"San Francisco International Airport 1948~ (campus)"
"SFO Terminal Complex 2017~ (building)"
"SFO Terminal Complex 2020-~05 (building)"
"SFO Terminal Complex 2021-05-25 (building)"
"SFO Terminal Complex 2021-11-09 (building)"
"California null (region)"
"SFO Terminal Complex 2006~ (building)"
"SFO Terminal Complex 2017~ (building)"
"SFO Terminal Complex 2020-~05 (building)"
"SFO Terminal Complex 2021-05-25 (building)"
"SFO Terminal Complex 2021-11-09 (building)"
"San Francisco-Oakland-San Jose null (marketarea)"
"SFO Terminal Complex 2000~ (building)"
"SFO Terminal Complex 2006~ (building)"
"SFO Terminal Complex 2017~ (building)"
"SFO Terminal Complex 2019-07-23 (building)"
"SFO Terminal Complex 2020-~05 (building)"
"SFO Terminal Complex 2021-05-25 (building)"
"SFO Terminal Complex 2021-11-09 (building)"
"San Mateo null (county)"
"SFO Terminal Complex 2000~ (building)"
"SFO Terminal Complex 2006~ (building)"
"SFO Terminal Complex 2011~ (building)"
"SFO Terminal Complex 2014~ (building)"
"SFO Terminal Complex 2017~ (building)"
"SFO Terminal Complex 2019-07-23 (building)"
"SFO Terminal Complex 2020-~05 (building)"
"SFO Terminal Complex 2021-05-25 (building)"
"SFO Terminal Complex 2021-11-09 (building)"
"International Terminal 2000~ (wing)"
"International Terminal 2006~ (wing)"
"International Terminal 2011~ (wing)"
"International Terminal 2014~ (wing)"
"International Terminal 2017~ (wing)"
"International Terminal 2019-07-23 (wing)"
"International Terminal 2020-~05 (wing)"
"International Terminal 2021-05-25 (wing)"
"International Terminal 2021-11-09 (wing)"
"International Terminal Main Hall 2017~ (concourse)"
"International Terminal Main Hall 2019-07-23 (concourse)"
"International Terminal Main Hall 2020-~05 (concourse)"
"International Terminal Main Hall 2021-05-25 (concourse)"
"International Terminal Main Hall 2021-11-09 (concourse)"
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial-grpc
* https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles
* https://github.com/protomaps