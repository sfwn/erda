package jsonpath

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestJQ(t *testing.T) {
	jsonInput := `{"success":true,"data":[{"instanceId":"ca3f094f903c14f5cb91943ade1113a9e","name":"应用监控","tag":"","addonName":"monitor","displayName":"应用监控","desc":"微服务监控可以洞察完整事务流程的性能表现，从前端APP或浏览器到后端服务器、容器和代码各个级别进行应用性能分析，对系统的可用性和性能进行可视化管理，支持实时告警和全栈根因定位，提升运维开发效率。","logoUrl":"//terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2019/10/15/9660d077-1404-4c18-974c-a18abc57569e.png","plan":"professional","version":"3.6","category":"monitoring&logging","config":{"PUBLIC_HOST":"","TERMINUS_AGENT_ENABLE":"","TERMINUS_KEY":"","TERMINUS_TA_COLLECTOR_URL":"","TERMINUS_TA_ENABLE":"","TERMINUS_TA_URL":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"STAGING","status":"ATTACHED","realInstanceId":"l1f4df54ae359453aa9633a767ba5ddd8","reference":3,"attachCount":3,"platform":true,"platformServiceType":1,"canDel":false,"terminusKey":"l1f4df54ae359453aa9633a767ba5ddd8","consoleUrl":"{\"tenantId\":\"l1f4df54ae359453aa9633a767ba5ddd8\",\"tenantGroup\":\"072ee655d8889c822db82313d83de5a9\",\"terminusKey\":\"l1f4df54ae359453aa9633a767ba5ddd8\",\"key\":\"Overview\"}","createdAt":"2020-11-16T17:30:24+08:00","updatedAt":"2020-11-20T18:47:54+08:00","recordId":0,"customAddonType":""},{"instanceId":"d87c3a79575d24c07a0f110040ab99732","name":"mysql","tag":"","addonName":"mysql","displayName":"MySQL","desc":"MySQL 是一种快速易用的 RDBMS，很多企业（不分规模大小）都在使用它来构建自己的数据库。MySQL 由一家瑞典公司 MySQL AB 开发、运营并予以支持。","logoUrl":"//terminus-dice.oss-cn-hangzhou.aliyuncs.com/addon/ui/icon/mysql.png","plan":"basic","version":"5.7.23","category":"database","config":{"ADDON_HAS_ENCRIPY":"","MYSQL_HOST":"","MYSQL_PASSWORD":"","MYSQL_PORT":"","MYSQL_SLAVE_HOST":"","MYSQL_SLAVE_PORT":"","MYSQL_USERNAME":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"TEST","status":"ATTACHED","realInstanceId":"w0e448ba11f8c46eeb8045834e4e02204","reference":1,"attachCount":1,"platform":false,"platformServiceType":0,"canDel":false,"consoleUrl":"","createdAt":"2020-11-16T10:19:17+08:00","updatedAt":"2020-11-16T10:22:05+08:00","recordId":0,"customAddonType":""},{"instanceId":"hffef9fa866c740538e88b5a89b790bf0","name":"日志分析","tag":"","addonName":"log-analytics","displayName":"日志分析","desc":"日志分析服务能够收集应用的日志进行存储，并提供全文搜索和日志统计等功能","logoUrl":"//terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2019/10/15/b2a6af21-1673-4fc8-a8c4-30d396f81e7e.png","plan":"basic","version":"2.0.0","category":"general_ability","config":{"MONITOR_LOG_COLLECTOR":"","MONITOR_LOG_KEY":"","MONITOR_LOG_OUTPUT":"","MONITOR_LOG_OUTPUT_CONFIG":"","TERMINUS_LOG_KEY":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"TEST","status":"ATTACHED","realInstanceId":"tab43de2040794e32a76ffe0c03c7aaa2","reference":1,"attachCount":1,"platform":true,"platformServiceType":2,"canDel":false,"consoleUrl":"","createdAt":"2020-11-16T10:19:17+08:00","updatedAt":"2020-11-16T10:27:46+08:00","recordId":0,"customAddonType":""},{"instanceId":"i317499b843ad4327bbb85710ec80b660","name":"应用监控","tag":"","addonName":"monitor","displayName":"应用监控","desc":"微服务监控可以洞察完整事务流程的性能表现，从前端APP或浏览器到后端服务器、容器和代码各个级别进行应用性能分析，对系统的可用性和性能进行可视化管理，支持实时告警和全栈根因定位，提升运维开发效率。","logoUrl":"//terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2019/10/15/9660d077-1404-4c18-974c-a18abc57569e.png","plan":"professional","version":"3.6","category":"monitoring&logging","config":{"PUBLIC_HOST":"","TERMINUS_AGENT_ENABLE":"","TERMINUS_KEY":"","TERMINUS_TA_COLLECTOR_URL":"","TERMINUS_TA_ENABLE":"","TERMINUS_TA_URL":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"DEV","status":"ATTACHED","realInstanceId":"vd41d9cc7ec28490dab00fbb136378cef","reference":1,"attachCount":1,"platform":true,"platformServiceType":1,"canDel":false,"terminusKey":"vd41d9cc7ec28490dab00fbb136378cef","consoleUrl":"{\"tenantId\":\"vd41d9cc7ec28490dab00fbb136378cef\",\"tenantGroup\":\"c8fde3376155a960d7ef87a50907083b\",\"terminusKey\":\"vd41d9cc7ec28490dab00fbb136378cef\",\"key\":\"Overview\"}","createdAt":"2020-11-16T11:18:01+08:00","updatedAt":"2020-11-27T15:48:38+08:00","recordId":0,"customAddonType":""},{"instanceId":"l6e6fcab6e37d43d6aa9a379307869754","name":"api-custom","tag":"","addonName":"custom","displayName":"Custom","desc":"提供第三方应用环境变量的配置。","logoUrl":"//terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2018/06/26/56d86431-0d22-4eb2-8111-30efcfc65a26.jpeg","plan":"basic","version":"1.0.0","category":"custom","config":{"key1":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"TEST","status":"ATTACHED","realInstanceId":"a7727546d034b45e7b0a0587455e3be1d","reference":0,"attachCount":0,"platform":false,"platformServiceType":0,"canDel":false,"consoleUrl":"","createdAt":"2020-12-02T17:56:42+08:00","updatedAt":"2020-12-02T17:56:42+08:00","recordId":0,"customAddonType":""},{"instanceId":"l9d7a376d5b7840039677db84e8cd9277","name":"redis","tag":"","addonName":"redis","displayName":"Redis","desc":"Redis是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的API。","logoUrl":"//terminus-dice.oss-cn-hangzhou.aliyuncs.com/addon/ui/icon/redis.png","plan":"basic","version":"3.2.12","category":"database","config":{"REDIS_HOST":"","REDIS_PASSWORD":"","REDIS_PORT":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"TEST","status":"ATTACHED","realInstanceId":"bb7edb8ae79954937929482c760762d29","reference":2,"attachCount":2,"platform":false,"platformServiceType":0,"canDel":false,"consoleUrl":"","createdAt":"2020-11-16T10:19:18+08:00","updatedAt":"2020-11-18T13:36:44+08:00","recordId":0,"customAddonType":""},{"instanceId":"n668c53eaeff1413884bd2b0deed8e50f","name":"应用监控","tag":"","addonName":"monitor","displayName":"应用监控","desc":"微服务监控可以洞察完整事务流程的性能表现，从前端APP或浏览器到后端服务器、容器和代码各个级别进行应用性能分析，对系统的可用性和性能进行可视化管理，支持实时告警和全栈根因定位，提升运维开发效率。","logoUrl":"//terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2019/10/15/9660d077-1404-4c18-974c-a18abc57569e.png","plan":"professional","version":"3.6","category":"monitoring&logging","config":{"PUBLIC_HOST":"","TERMINUS_AGENT_ENABLE":"","TERMINUS_KEY":"","TERMINUS_TA_COLLECTOR_URL":"","TERMINUS_TA_ENABLE":"","TERMINUS_TA_URL":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"TEST","status":"ATTACHED","realInstanceId":"ob631a0081d4640c5bae28e8f72d96c1a","reference":2,"attachCount":2,"platform":true,"platformServiceType":1,"canDel":false,"terminusKey":"ob631a0081d4640c5bae28e8f72d96c1a","consoleUrl":"{\"tenantId\":\"ob631a0081d4640c5bae28e8f72d96c1a\",\"tenantGroup\":\"d3905dc1f906d71217393e4ae5020071\",\"terminusKey\":\"ob631a0081d4640c5bae28e8f72d96c1a\",\"key\":\"Overview\"}","createdAt":"2020-11-16T10:19:18+08:00","updatedAt":"2020-11-18T13:36:44+08:00","recordId":0,"customAddonType":""},{"instanceId":"o2877235e271d476cad4d96dc20247586","name":"应用监控","tag":"","addonName":"monitor","displayName":"应用监控","desc":"微服务监控可以洞察完整事务流程的性能表现，从前端APP或浏览器到后端服务器、容器和代码各个级别进行应用性能分析，对系统的可用性和性能进行可视化管理，支持实时告警和全栈根因定位，提升运维开发效率。","logoUrl":"//terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2019/10/15/9660d077-1404-4c18-974c-a18abc57569e.png","plan":"professional","version":"3.6","category":"monitoring&logging","config":{"PUBLIC_HOST":"","TERMINUS_AGENT_ENABLE":"","TERMINUS_KEY":"","TERMINUS_TA_COLLECTOR_URL":"","TERMINUS_TA_ENABLE":"","TERMINUS_TA_URL":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"PROD","status":"ATTACHED","realInstanceId":"cd0d37ed0041b49a2a13901a3bb4cea31","reference":2,"attachCount":2,"platform":true,"platformServiceType":1,"canDel":false,"terminusKey":"cd0d37ed0041b49a2a13901a3bb4cea31","consoleUrl":"{\"tenantId\":\"cd0d37ed0041b49a2a13901a3bb4cea31\",\"tenantGroup\":\"d3a75d30e812c606cf04b870b7a539ab\",\"terminusKey\":\"cd0d37ed0041b49a2a13901a3bb4cea31\",\"key\":\"Overview\"}","createdAt":"2020-11-18T11:40:55+08:00","updatedAt":"2020-11-18T11:51:49+08:00","recordId":0,"customAddonType":""},{"instanceId":"q3fd8763739564d1da1b2d8c9c9692929","name":"redis","tag":"","addonName":"redis","displayName":"Redis","desc":"Redis是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的API。","logoUrl":"//terminus-dice.oss-cn-hangzhou.aliyuncs.com/addon/ui/icon/redis.png","plan":"basic","version":"3.2.12","category":"database","config":{"REDIS_HOST":"","REDIS_PASSWORD":"","REDIS_PORT":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"PROD","status":"ATTACHED","realInstanceId":"mf6c174a7a7794c40bd4177b8f829caf2","reference":2,"attachCount":2,"platform":false,"platformServiceType":0,"canDel":false,"consoleUrl":"","createdAt":"2020-11-18T11:40:55+08:00","updatedAt":"2020-11-18T11:51:48+08:00","recordId":0,"customAddonType":""},{"instanceId":"s7541918066ac4a3ca0c149feb45d1eba","name":"2","tag":"","addonName":"custom","displayName":"Custom","desc":"提供第三方应用环境变量的配置。","logoUrl":"//terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2018/06/26/56d86431-0d22-4eb2-8111-30efcfc65a26.jpeg","plan":"basic","version":"1.0.0","category":"custom","config":{"key2":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"TEST","status":"ATTACHED","realInstanceId":"ufe99f555d7dd4821ad07e83aa720dd2b","reference":0,"attachCount":0,"platform":false,"platformServiceType":0,"canDel":false,"consoleUrl":"","createdAt":"2020-12-02T17:42:38+08:00","updatedAt":"2020-12-02T17:42:38+08:00","recordId":0,"customAddonType":""},{"instanceId":"x8877c9e45aa343a6be6c7219960c097d","name":"redis","tag":"","addonName":"redis","displayName":"Redis","desc":"Redis是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的API。","logoUrl":"//terminus-dice.oss-cn-hangzhou.aliyuncs.com/addon/ui/icon/redis.png","plan":"basic","version":"3.2.12","category":"database","config":{"REDIS_HOST":"","REDIS_PASSWORD":"","REDIS_PORT":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"DEV","status":"ATTACHED","realInstanceId":"y913c245f27554cbe9bfe9350658344f7","reference":0,"attachCount":0,"platform":false,"platformServiceType":0,"canDel":false,"consoleUrl":"","createdAt":"2020-11-16T11:18:01+08:00","updatedAt":"2020-11-18T13:37:13+08:00","recordId":0,"customAddonType":""},{"instanceId":"z2ed4cddbf27a41efae24641b9b478af8","name":"redis","tag":"","addonName":"redis","displayName":"Redis","desc":"Redis是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的API。","logoUrl":"//terminus-dice.oss-cn-hangzhou.aliyuncs.com/addon/ui/icon/redis.png","plan":"basic","version":"3.2.12","category":"database","config":{"REDIS_HOST":"","REDIS_PASSWORD":"","REDIS_PORT":""},"shareScope":"PROJECT","cluster":"terminus-test","orgId":3,"projectId":2,"projectName":"ss_pro1","workspace":"STAGING","status":"ATTACHED","realInstanceId":"g5652ad6b121f4e9d8d314a2377433fff","reference":3,"attachCount":3,"platform":false,"platformServiceType":0,"canDel":false,"consoleUrl":"","createdAt":"2020-11-16T17:30:24+08:00","updatedAt":"2020-11-20T18:47:54+08:00","recordId":0,"customAddonType":""}],"err":{"code":"","msg":"","ctx":null}}`
	result, err := JQ(jsonInput, `.data`)
	assert.NoError(t, err)
	spew.Dump(result)
}

func TestJQ2(t *testing.T) {
	jsonInput := "{\"sessionid\":\"aebaac25-fd1b-476f-a79b-0e10df2ad8c2\",\"id\":2,\"token\":\"\",\"email\":\"dice@dice.terminus.io\",\"emailExist\":true,\"passwordExist\":false,\"phoneExist\":false,\"birthday\":\"\",\"passwordStrength\":0,\"phone\":\"\",\"avatarUrl\":\"\",\"username\":\"dice\",\"nickName\":\"dice\",\"enabled\":true,\"createdAt\":\"2019-08-23 04:15:15\",\"updatedAt\":\"\",\"lastLoginAt\":\"\"}"

	result, err := JQ(jsonInput, `.id`)
	assert.NoError(t, err)
	spew.Dump(result)

	result, err = JQ(jsonInput, `.sessionid`)
	assert.NoError(t, err)
	spew.Dump(result)
}
