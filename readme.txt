#MSvrs
1，c/s`s server frame,with golang design
2，client <-[1]-> Gate <-[2]-> Svrs
	[1]RPC调用
	[2]redis-mq:pub/sub
3, db storage: redis,Mysql
4, 带宽成本控制：
	Gate配置外网地址；redis／Svrs配置内网地址。
5，QPS：
6，TPS：	
