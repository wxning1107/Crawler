# Golang_crawler

Crawler is a distributed web crawler written in golang without using any crawler framework.


It is a personal project :smile:, starting from 0 using the native code to build a distributed crawler system.


The main purpose is to deeply understand the concurrency mechanism of golang and the design idea of the distributed system.

[单任务版爬虫](https://wxning1107.github.io/2019/07/12/go-crawler1/)   
[并发版爬虫](https://wxning1107.github.io/2019/08/03/go-crawler2/)

# Features

- The breadth-first algorithm framework,embedded data crawling and the information extraction is applied to implement the basic crawler task.

- Utilize the natural advantages of Go in concurrency to achieve the distribution and scheduling of crawler tasks to achieve concurrent requirements.

- Using rpc to separate and be independent of concurrent tasks in a  single task version to implement distributed crawlers.

- Using Docker+ElasticSearch to build a data storage backend, using the Go template library for data display

<!-- ALL-CONTRIBUTORS-LIST: START - Do not remove or modify this section -->
<!-- ALL-CONTRIBUTORS-LIST:END -->
