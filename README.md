# port-agent
port agent 

基于小米运维开源的open-falcon，端口监控专用agent。

##采集的metric列表：
  port.agent.alive
  port.alive
  
##配置说明
配置文件请参照cfg.example.json，修改该文件名为cfg.json，将该文件里的IP换成实际使用的IP。
```
{
    "debug": true,
    "hostname": "zsf-test-port",                           // endpoint
    "ip": "192.168.6.9",                                   //需要被探测的ip(域名)
    "plugin": {
        "enabled": false,
        "dir": "./plugin",
        "git": "https://github.com/open-falcon/plugin.git",
        "logs": "./logs"
    },
    "heartbeat": {
        "enabled": true,
        "addr": "127.0.0.1:6030",
        "interval": 60,
        "timeout": 1000
    },
    "transfer": {
        "enabled": true,
        "addrs": [
            "127.0.0.1:8433",
            "127.0.0.1:8433"
        ],
        "interval": 60,
        "timeout": 1000
    },
    "http": {
        "enabled": true,
        "listen": ":1988",
        "backdoor": false
    },
    "port": ["5678"]                                            // 指定IP或域名需要被探测的端口列表

}
```

