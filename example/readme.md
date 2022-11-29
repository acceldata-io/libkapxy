# TinyKap

Example tool that extracts the Kafka producer names and the topic which they're producing to

---

Compile `tinykap` binary for Linux:

```shell
go build -a -gcflags=all="-l -B" -ldflags "-s -w" .
```

Output:

```shell
2022-11-22 13:19:06.19609 +0530 IST   map[apiVersion:5 host:10.90.5.55 messageSize:170 producer:store_ProdCount_0]   map[producer:store_ProdCount_0 topic:store]
2022-11-22 13:19:06.635998 +0530 IST   map[apiVersion:5 host:10.90.5.55 messageSize:192 producer:customer_address_ProdCount_0]   map[producer:customer_address_ProdCount_0 topic:customer_address]
2022-11-22 13:19:06.636281 +0530 IST   map[apiVersion:5 host:10.90.5.55 messageSize:183 producer:dummyTopic1_ProdCount_0]   map[producer:dummyTopic1_ProdCount_0 topic:dummyTopic1]
2022-11-22 13:19:06.636388 +0530 IST   map[apiVersion:5 host:10.90.5.55 messageSize:182 producer:dummyTopic2_ProdCount_0]   map[producer:dummyTopic2_ProdCount_0 topic:dummyTopic2]
2022-11-22 13:19:06.651171 +0530 IST   map[apiVersion:5 host:10.90.5.55 messageSize:187 producer:manyPartition_ProdCount_0]   map[producer:manyPartition_ProdCount_0 topic:manyPartition]
2022-11-22 13:19:06.696256 +0530 IST   map[apiVersion:5 host:10.90.5.55 messageSize:170 producer:store_ProdCount_0]   map[producer:store_ProdCount_0 topic:store]
```

---
