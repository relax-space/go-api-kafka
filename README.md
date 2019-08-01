# go-api-producer

## prepare

install kafka
```
https://www.cnblogs.com/mignet/p/window_kafka.html
https://www.apache.org/dyn/closer.cgi?path=/kafka/2.3.0/kafka_2.12-2.3.0.tgz
```

## send message
http://localhost:8090/send

## accept message

```bash
$ ./kafka-console-consumer.bat --bootstrap-server localhost:9092 --from-beginning --topic fruit
 {"authToken":"","createdAt":"2019-07-30T23:51:30.5899707Z","payload":{"name":"apple"},"requestId":"","status":"FruitCreated"}
```