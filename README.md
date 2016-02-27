sanic
=====

sanic is a clone of
[Twitter snowflake](https://github.com/twitter/snowflake/tree/snowflake-2010)
(the 2010 version), written in [Golang](https://golang.org/).
More specifically, the [IdWorker section of snowflake]
(https://github.com/twitter/snowflake/blob/snowflake-2010/src/main/scala/com/twitter/service/snowflake/IdWorker.scala).

Check out [the examples](https://github.com/ifo/sanic/tree/master/examples) for
how to use it.

Currently it only generates 10 character ids, though future updates will likely
make that configurable.

## License

sanic is ISC licensed.
Check out the [LICENSE](https://github.com/ifo/sanic/blob/master/LICENSE) file.
