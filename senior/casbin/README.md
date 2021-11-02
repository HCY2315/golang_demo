## casbin 表中个键的含义
博客链接：https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/%E5%85%B6%E4%BB%96/%E6%9D%83%E9%99%90%E7%AE%A1%E7%90%86.html
```
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL COMMENT '类型',
  `v0` varchar(100) DEFAULT NULL COMMENT '角色',
  `v1` varchar(100) DEFAULT NULL COMMENT '路径',
  `v2` varchar(100) DEFAULT NULL COMMENT '请求方法',
  `v3` varchar(100) DEFAULT NULL COMMENT '允许读写',
  `v4` varchar(100) DEFAULT NULL COMMENT '允许/拒绝',
  `v5` varchar(100) DEFAULT NULL,
  KEY `IDX_casbin_rule_v4` (`v4`),
  KEY `IDX_casbin_rule_v5` (`v5`),
  KEY `IDX_casbin_rule_p_type` (`p_type`),
  KEY `IDX_casbin_rule_v0` (`v0`),
  KEY `IDX_casbin_rule_v1` (`v1`),
  KEY `IDX_casbin_rule_v2` (`v2`),
  KEY `IDX_casbin_rule_v3` (`v3`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
```
casbin_rule表中
  p_type --类型，可以是p策略，g角色等等
  v0 --角色 roleName/roleId   sub
  v1 --Path 路径              obj
  v2 --Method 请求方式         act
  v3 --允许读/写 read/write
  v4 --允许/拒绝 allow/deny
  v5 --不知道

  ptype=g的时候v1=角色
  ptype=p的时候，v2=action，v0=subject，v1=obj
  策略：即ptype为p的数据
```