## Run
```shell
$ go run main.go \
  --name node1 \
  --port 8080
  
$ go run main.go \
  --name node2 \
  --port 8081 \
  --join localhost:8080
  
 $ go run main.go \
  --name node3 \
  --port 8082 \
  --join localhost:8080
```

## Output
### Node 1
```shell
2023/02/13 18:04:08 Node at 172.17.224.215:8080:%!s(uint16=8080)
2023/02/13 18:04:08 Member: name1(172.17.224.215:8080)
2023/02/13 18:04:46 [DEBUG] memberlist: Stream connection from=127.0.0.1:38186
2023/02/13 18:04:51 [DEBUG] memberlist: Failed ping: node2 (timeout reached)
2023/02/13 18:04:52 [DEBUG] memberlist: Stream connection from=127.0.0.1:38190
2023/02/13 18:04:52 [INFO] memberlist: Suspect node2 has failed, no acks received
2023/02/13 18:04:59 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:05:01 [DEBUG] memberlist: Stream connection from=127.0.0.1:38194
2023/02/13 18:05:14 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:05:21 [DEBUG] memberlist: Stream connection from=172.17.224.215:47024
2023/02/13 18:05:29 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:05:36 [DEBUG] memberlist: Stream connection from=172.17.224.215:47028
2023/02/13 18:05:44 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:05:47 [DEBUG] memberlist: Stream connection from=127.0.0.1:38206
2023/02/13 18:05:59 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:06:14 [DEBUG] memberlist: Initiating push/pull sync with: node3 172.17.224.215:8083
2023/02/13 18:06:29 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:06:43 [DEBUG] memberlist: Stream connection from=172.17.224.215:47052
2023/02/13 18:06:44 [DEBUG] memberlist: Initiating push/pull sync with: node3 172.17.224.215:8083
2023/02/13 18:06:51 [DEBUG] memberlist: Stream connection from=172.17.224.215:47056
2023/02/13 18:06:59 [DEBUG] memberlist: Initiating push/pull sync with: node3 172.17.224.215:8083
```

### Node2 
```shell
2023/02/13 18:05:01 Node at 172.17.224.215:8081:%!s(uint16=8081)
2023/02/13 18:05:01 [DEBUG] memberlist: Initiating push/pull sync with:  127.0.0.1:8080
2023/02/13 18:05:01 [WARN] memberlist: Refuting an alive message for 'node2' (172.17.224.215:8081) meta:([] VS []), vsn:([1 5 2 0 0 0] VS [1 5 2 0 0 0])
2023/02/13 18:05:01 Member: name1(172.17.224.215:8080)
2023/02/13 18:05:01 Member: node2(172.17.224.215:8081)
2023/02/13 18:05:14 [DEBUG] memberlist: Stream connection from=172.17.224.215:51158
2023/02/13 18:05:21 [DEBUG] memberlist: Initiating push/pull sync with: name1 172.17.224.215:8080
2023/02/13 18:05:29 [DEBUG] memberlist: Stream connection from=172.17.224.215:51162
2023/02/13 18:05:36 [DEBUG] memberlist: Initiating push/pull sync with: name1 172.17.224.215:8080
2023/02/13 18:05:44 [DEBUG] memberlist: Stream connection from=172.17.224.215:51166
2023/02/13 18:05:51 [DEBUG] memberlist: Initiating push/pull sync with: node3 172.17.224.215:8083
2023/02/13 18:05:59 [DEBUG] memberlist: Stream connection from=172.17.224.215:51172
2023/02/13 18:06:06 [DEBUG] memberlist: Initiating push/pull sync with: node3 172.17.224.215:8083
2023/02/13 18:06:13 [DEBUG] memberlist: Stream connection from=172.17.224.215:51176
2023/02/13 18:06:21 [DEBUG] memberlist: Initiating push/pull sync with: node3 172.17.224.215:8083
```

### Node3
```shell
2023/02/13 18:05:47 Node at 172.17.224.215:8083:%!s(uint16=8083)
2023/02/13 18:05:47 [DEBUG] memberlist: Initiating push/pull sync with:  127.0.0.1:8080
2023/02/13 18:05:47 Member: node2(172.17.224.215:8081)
2023/02/13 18:05:47 Member: node3(172.17.224.215:8083)
2023/02/13 18:05:47 Member: name1(172.17.224.215:8080)
2023/02/13 18:05:51 [DEBUG] memberlist: Stream connection from=172.17.224.215:47860
2023/02/13 18:06:06 [DEBUG] memberlist: Stream connection from=172.17.224.215:47864
2023/02/13 18:06:13 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:06:14 [DEBUG] memberlist: Stream connection from=172.17.224.215:47868
2023/02/13 18:06:21 [DEBUG] memberlist: Stream connection from=172.17.224.215:47870
2023/02/13 18:06:28 [DEBUG] memberlist: Initiating push/pull sync with: node2 172.17.224.215:8081
2023/02/13 18:06:36 [DEBUG] memberlist: Stream connection from=172.17.224.215:47876
2023/02/13 18:06:43 [DEBUG] memberlist: Initiating push/pull sync with: name1 172.17.224.215:8080
```