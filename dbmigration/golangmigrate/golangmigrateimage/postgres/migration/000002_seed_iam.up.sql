INSERT INTO public.iam_orgs (id, created_at,created_by,updated_at,updated_by,"name",description) VALUES
	 (1,NOW(),NULL,NOW(),NULL,'admin','Default organization for administrators');

INSERT INTO public.iam_users (id, created_at,created_by,updated_at,updated_by,login,"password",name1,name2,name3,settings,has_password,has_api_key,"blocked",org_id) VALUES
	 (1,NOW(),NULL,NOW(),NULL,'admin','$2b$10$HD2k2y1hQG1iCxq6Snha2uDoWkUfp949EJiu4Br14VJ5rVgRby6Dy',NULL,NULL,NULL,'{"theme": {"darkMode": false}, "preferences": {}, "notifications": {}}',false,false,false,1);

INSERT INTO public.iam_roles (id, created_at,created_by,updated_at,updated_by,"name",description) VALUES
	 (1,NOW(),NULL,NOW(),NULL,'super-admin','Super Administrator with all permissions');

INSERT INTO public.iam_permissions (id, "action",subject,conditions,description,inverted) VALUES
	 (1,'manage','all',NULL,'Full access to all resources',false);

INSERT INTO public.iam_roles_permissions (role_id,permission_id) VALUES
	 (1,1);

INSERT INTO public.iam_users_roles (user_id,role_id) VALUES
	 (1,1);
