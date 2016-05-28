# bench-println-example
a example code used to identify a runtime issue of go while using fmt.Fprintln


#### 问题重现步骤
  * go run main.go  
    默认在9090端口开启两个endpoint:  localhost:9090/h1, localhost:9090/h2      
    h1使用fmt.Fprintf向控制台打印一串字符，h2使用log向控制台打印同样的字符串；
    
    
  * 使用wrk测试h1    
    wrk -t10 -d20 http://localhost:9090/h1  几秒钟后，服务异常终止；     
    异常信息为：    
    runtime/cgo: pthread_create failed: Resource temporarily unavailable     
    SIGABRT: abort    
    PC=0x7fff8c698f06 m=2   
    详细内容见stack_dump.log   
    
    测试期间，通过http://localhost:9090/debug/pprof可以观察到threadcreated剧增   
    
  * 使用wrk测试h2     
    wrk -t10 -d20 http://localhost:9090/h2  服务可以正常运行    
  
