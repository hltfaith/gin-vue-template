-- 创建数据库
create database manage_system_db;

-- 其他
drop table t_permission,t_role,t_role_permission,t_user,t_user_profile;

-- 模拟数据
-- 权限表
insert 
    into t_permission
        (permission_id,is_group,title,path,redirect,component,icon,sub_component,sub_path)
    values
        (1,'否','首页','/','/dashboard','Layout','dashboard','/views/dashboard/index','dashboard');

insert 
    into t_permission
        (permission_id,is_group,title,path,redirect,component,icon,sub_title,sub_component,sub_path,sub_icon)
    values
        (
         2,'是','权限管理','/permission','/permission','Layout','table','用户管理','/views/permission/user','user','user'
        );

insert 
    into t_permission
        (permission_id,is_group,title,path,redirect,component,icon,sub_title,sub_component,sub_path,sub_icon)
    values
        (
         3,'是','权限管理','/permission','/permission','Layout','table','角色管理','/views/permission/role','role','user'
        );

-- 角色表
insert into t_role(name,role_id,rid) values('管理员','admin',1);
insert into t_role(name,role_id,rid) values('普通用户','user',2);

-- 角色权限关联表
insert into t_role_permission(pid,rid) values(1, 1);
insert into t_role_permission(pid,rid) values(2, 1);
insert into t_role_permission(pid,rid) values(3, 1);
insert into t_role_permission(pid,rid) values(1, 2);

-- 用户表 默认密码 123456
insert into t_user(user_id,username,password,rid) values('10001','admin','bf1552ce87c50c6d8b43f11fe070a1a0',1);
insert into t_user(user_id,username,password,rid) values('10002','user1','bf1552ce87c50c6d8b43f11fe070a1a0',2);
