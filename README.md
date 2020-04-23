# go-api-producer

## start

### 1. run docker kafka
1.设置HOST文件：（xxx为你本机的ip，不能写成127.0.0.1）
    xxx test-kafka

2.运行以下命令：
```
$ docker-compose -f example/docker-compose.yml up -d
```

### 2. 运行go-api-consumer
```
go run .
```

## 3.send message
http://localhost:8090/send

## 4.accept message

```bash
$ ./kafka-console-consumer.bat --bootstrap-server localhost:9092 --from-beginning --topic fruit
 {"authToken":"","createdAt":"2019-07-30T23:51:30.5899707Z","payload":{"name":"apple"},"requestId":"","status":"FruitCreated"}
```

参考：
https://www.cnblogs.com/mignet/p/window_kafka.html
