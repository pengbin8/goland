# 1. netcore系列
netcore3.1
net5
net6
# 2. [ netcore3.1 &net6 ]注册到nacos，nacos配置管理，服务注册
# 3. [ CAP ] 解决分布式事务
# 4. [ 网关 ]
# 4.1 ocelot 网关，集成swagger  MMLib.SwaggerForOcelot （https://blog.51cto.com/u_13746169/5876532）
# 4.2 ocelot 网关，搭配cusul MMLib.SwaggerForOcelot （https://blog.51cto.com/u_13746169/5876532）
# 5. 领域设计 工作单元unitwork 解决事务一致性
# 6. 领域设计 领域事件，解耦微服务逻辑拆分，例如通知事件，日志事件，库存扣减事件，将原来在一个方法内的逻辑拆分成多个领域，多个方法来执行，而且保证多个方法执行的原子性和一致性
# 7. 领域设计 集成事件，是领域事件的扩充，将多个方法直接拆分成多个微服务，集成事件就是保证多个微服务执行的原子性和一致性
# 8.WebApiClientCore的使用 【HttpGet/HttpPost】 API接口处理

