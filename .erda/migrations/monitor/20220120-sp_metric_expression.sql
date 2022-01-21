UPDATE `sp_metric_expression` SET `enable` = 0 WHERE `id` in (SELECT `id` FROM (SELECT `id` FROM `sp_metric_expression` WHERE `enable` = 1 AND `expression` LIKE "%\"alias\":\"application_cache_service\"%") AS `ids`);
UPDATE `sp_metric_expression` SET `enable` = 0 WHERE `id` in (SELECT `id` FROM (SELECT `id` FROM `sp_metric_expression` WHERE `enable` = 1 AND `expression` LIKE "%\"alias\":\"application_db_service\"%") AS `ids`);
UPDATE `sp_metric_expression` SET `enable` = 0 WHERE `id` in (SELECT `id` FROM (SELECT `id` FROM `sp_metric_expression` WHERE `enable` = 1 AND `expression` LIKE "%\"alias\":\"application_mq_service\"%") AS `ids`);

INSERT `sp_metric_expression`(`attributes`,`expression`,`version`) VALUES("{}","{\"alias\":\"application_cache_service\",\"filter\":{},\"functions\":[{\"aggregator\":\"sum\",\"field\":\"elapsed\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_mean * field.elapsed_count; }\"},{\"aggregator\":\"sum\",\"field\":\"count\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_count; }\"},{\"aggregator\":\"sum\",\"field\":\"errors\",\"field_script\":\"function invoke(field, tag){ return tag.error=='true'?field.elapsed_count:0; }\"},{\"aggregator\":\"rateps\",\"field\":\"reqs_per_second\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_count; }\"}],\"group\":[\"component\",\"host\",\"source_terminus_key\",\"source_application_id\",\"source_runtime_name\",\"source_service_name\",\"source_addon_id\",\"error\"],\"metric\":\"application_cache\",\"outputs\":[\"metric\"],\"select\":{\"_metric_scope\":\"#_metric_scope\",\"_metric_scope_id\":\"#_metric_scope_id\",\"cluster_name\":\"#cluster_name\",\"component\":\"#component\",\"db_host\":\"#db_host\",\"db_system\":\"#db_system\",\"host\":\"#host\",\"source_addon_id\":\"#source_addon_id\",\"source_addon_type\":\"#source_addon_type\",\"source_application_id\":\"#source_application_id\",\"source_application_name\":\"#source_application_name\",\"source_project_id\":\"#source_project_id\",\"source_project_name\":\"#source_project_name\",\"source_runtime_id\":\"#source_runtime_id\",\"source_runtime_name\":\"#source_runtime_name\",\"source_service_id\":\"#source_service_id\",\"source_service_name\":\"#source_service_name\",\"source_terminus_key\":\"#source_terminus_key\",\"source_workspace\":\"#source_workspace\"},\"window\":1}","3.0");
INSERT `sp_metric_expression`(`attributes`,`expression`,`version`) VALUES("{}","{\"alias\":\"application_db_service\",\"filter\":{},\"functions\":[{\"aggregator\":\"sum\",\"field\":\"elapsed\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_mean * field.elapsed_count; }\"},{\"aggregator\":\"sum\",\"field\":\"count\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_count; }\"},{\"aggregator\":\"sum\",\"field\":\"errors\",\"field_script\":\"function invoke(field, tag){ return tag.error=='true'?field.elapsed_count:0; }\"},{\"aggregator\":\"rateps\",\"field\":\"reqs_per_second\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_count; }\"}],\"group\":[\"component\",\"host\",\"source_terminus_key\",\"source_application_id\",\"source_runtime_name\",\"source_service_name\",\"source_addon_id\",\"error\"],\"metric\":\"application_db\",\"outputs\":[\"metric\"],\"select\":{\"cluster_name\":\"#cluster_name\",\"component\":\"#component\",\"db_host\":\"#db_host\",\"db_system\":\"#db_system\",\"host\":\"#host\",\"source_addon_id\":\"#source_addon_id\",\"source_addon_type\":\"#source_addon_type\",\"source_application_id\":\"#source_application_id\",\"source_application_name\":\"#source_application_name\",\"source_project_id\":\"#source_project_id\",\"source_project_name\":\"#source_project_name\",\"source_runtime_id\":\"#source_runtime_id\",\"source_runtime_name\":\"#source_runtime_name\",\"source_service_id\":\"#source_service_id\",\"source_service_name\":\"#source_service_name\",\"source_terminus_key\":\"#source_terminus_key\",\"source_workspace\":\"#source_workspace\"},\"window\":1}","3.0");
INSERT `sp_metric_expression`(`attributes`,`expression`,`version`) VALUES("{}","{\"alias\":\"application_mq_service\",\"filter\":{},\"functions\":[{\"aggregator\":\"sum\",\"field\":\"elapsed\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_mean * field.elapsed_count; }\"},{\"aggregator\":\"sum\",\"field\":\"count\",\"field_script\":\"function invoke(field, tag){ return field.elapsed_count; }\"}],\"group\":[\"component\",\"host\",\"target_terminus_key\",\"target_application_id\",\"target_runtime_name\",\"target_service_name\",\"target_addon_id\",\"source_terminus_key\",\"source_application_id\",\"source_runtime_name\",\"source_service_name\",\"source_addon_id\"],\"metric\":\"application_mq\",\"outputs\":[\"metric\"],\"select\":{\"cluster_name\":\"#cluster_name\",\"component\":\"#component\",\"db_type\":\"#db_type\",\"host\":\"#host\",\"peer_address\":\"#peer_address\",\"peer_service\":\"#peer_service\",\"source_addon_id\":\"#source_addon_id\",\"source_addon_type\":\"#source_addon_type\",\"source_application_id\":\"#source_application_id\",\"source_application_name\":\"#source_application_name\",\"source_project_id\":\"#source_project_id\",\"source_project_name\":\"#source_project_name\",\"source_runtime_id\":\"#source_runtime_id\",\"source_runtime_name\":\"#source_runtime_name\",\"source_service_id\":\"#source_service_id\",\"source_service_name\":\"#source_service_name\",\"source_terminus_key\":\"#source_terminus_key\",\"source_workspace\":\"#source_workspace\",\"target_addon_id\":\"#target_addon_id\",\"target_addon_type\":\"#target_addon_type\",\"target_application_id\":\"#target_application_id\",\"target_application_name\":\"#target_application_name\",\"target_instance_id\":\"#target_instance_id\",\"target_project_id\":\"#target_project_id\",\"target_project_name\":\"#target_project_name\",\"target_runtime_name\":\"#target_runtime_name\",\"target_service_id\":\"#target_service_id\",\"target_service_name\":\"#target_service_name\",\"target_terminus_key\":\"#target_terminus_key\",\"target_workspace\":\"#target_workspace\"},\"window\":3}","3.0");
