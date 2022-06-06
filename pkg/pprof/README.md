## pprof 基础

### 查看当前内存使用
```
> go tool pprof http://localhost:8007/debug/pprof/heap
Fetching profile over HTTP from http://localhost:8007/debug/pprof/heap
Saved profile in /root/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
Type: inuse_space
Time: Jun 6, 2022 at 10:17pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

### 采集30秒 cpu 性能信息
```
> go tool pprof http://localhost:8007/debug/pprof/profile\?seconds\=30
Fetching profile over HTTP from http://localhost:8007/debug/pprof/profile?seconds=30
Saved profile in /root/pprof/pprof.samples.cpu.001.pb.gz
Type: cpu
Time: Jun 6, 2022 at 10:19pm (CST)
Duration: 30s, Total samples = 90ms (  0.3%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

### 使用 web 界面查看结果
```
> go tool pprof -http=:9999 /root/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
> go tool pprof -http=:9999 /root/pprof/pprof.samples.cpu.001.pb.gz
```

在浏览器中打开地址：http://localhost:9999/

## 进一步阅读

* [Profiling Go Programs](https://go.dev/blog/pprof)