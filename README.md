Modbeat is an example Community beat. It collects the file information for given globs on a predefined period and sends the follwoing info about each file to Elasticsearch or Logstash:

* name: File name
* path: Absolute file path including file name
* mode:
* modtime: Last modification time
* size: Size of the file

An example JSON document stored in elasticsearch looks as following:

```
{
   "@timestamp":"2016-01-19T23:48:20.423Z",
   "beat":{
      "hostname":"examplehost",
      "name":"ruflin"
   },
   "count":1,
   "mode":420,
   "modtime":"2016-01-16T00:10:33+01:00",
   "name":"example.log",
   "path":"/var/log/example.log",
   "size":3752,
   "type":"modbeat"
}
```


# Goal

Modbeat has the following goals:

* Basic beat example
* Foundation for maintaining and updating beat
* Structure to create and maintain community around the beat
* Ensure quality of the beat (testing, structure)
* Foundation for good beat docs


# TODO

* Which LICENSE do we recommend? Apache / MIT?
* Make system-tests should not run tests for vendor packages
* Fix System tests
* Complete CONTRIBUTING.md
* Document config

# Notes

* Copy files (already with vendor inside?)
* Basic doc file?
* Glide init, glide update
* Go vendor magic needed to not gofmt all libs: https://gist.github.com/bgentry/fd1ffef7dbde01857f66
* Move modbeat.yml in top directory -> direct start of binary is possible
* Tag libbeat to specific version tag?
