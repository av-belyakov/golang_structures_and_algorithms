-- Drop all tables in reverse order of dependencies

-- Drop views first
DROP TABLE IF EXISTS public.users_roles CASCADE;
DROP TABLE IF EXISTS public.users_metrics CASCADE;
DROP TABLE IF EXISTS public.users_histories CASCADE;
DROP TABLE IF EXISTS public.users_actions CASCADE;
DROP TABLE IF EXISTS public.users CASCADE;

-- Drop statistics and other standalone tables
DROP TABLE IF EXISTS public.statistics CASCADE;

-- Drop socbox tables
DROP TABLE IF EXISTS public.socbox_geo_regions CASCADE;
DROP TABLE IF EXISTS public.socbox_geo_districts CASCADE;
DROP TABLE IF EXISTS public.socbox_geo_countries CASCADE;
DROP TABLE IF EXISTS public.socbox_geo_cities CASCADE;

DROP TABLE IF EXISTS public.socbox_extras_tagged_item CASCADE;
DROP TABLE IF EXISTS public.socbox_extras_tag_object_types CASCADE;
DROP TABLE IF EXISTS public.socbox_extras_tags CASCADE;
DROP TABLE IF EXISTS public.socbox_extras_menu CASCADE;
DROP TABLE IF EXISTS public.socbox_extras_custom_fields_objects_types CASCADE;
DROP TABLE IF EXISTS public.socbox_extras_custom_fields_choicesets CASCADE;
DROP TABLE IF EXISTS public.socbox_extras_custom_fields CASCADE;
DROP TABLE IF EXISTS public.socbox_extras_changelog CASCADE;

DROP TABLE IF EXISTS public.socbox_detection_threat_levels CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_signatures_classtypes CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_signatures CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_sensors_hardwares CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_sensors_groups CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_sensors CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_recomendations CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_nccci_classes CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_mitre_techniques CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_mitre_tactics CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_local_classtypes CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_directions CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_classtypes CASCADE;
DROP TABLE IF EXISTS public.socbox_detection_attacks_priorities CASCADE;

DROP TABLE IF EXISTS public.socbox_core_content_type CASCADE;
DROP TABLE IF EXISTS public.socbox_content_type CASCADE;

DROP TABLE IF EXISTS public.socbox_assets_resources CASCADE;
DROP TABLE IF EXISTS public.socbox_assets_resources_types CASCADE;
DROP TABLE IF EXISTS public.socbox_assets_homenets CASCADE;

DROP TABLE IF EXISTS public.socbox_actors_subjects CASCADE;
DROP TABLE IF EXISTS public.socbox_actors_scopes CASCADE;
DROP TABLE IF EXISTS public.socbox_actors_persons_orders CASCADE;
DROP TABLE IF EXISTS public.socbox_actors_objects_sensors CASCADE;
DROP TABLE IF EXISTS public.socbox_actors_objects CASCADE;

-- Drop SIEM tables
DROP TABLE IF EXISTS public.siem_views CASCADE;
DROP TABLE IF EXISTS public.siem_tasks CASCADE;
DROP TABLE IF EXISTS public.siem_task_types CASCADE;
DROP TABLE IF EXISTS public.siem_task_statuses CASCADE;
DROP TABLE IF EXISTS public.siem_filters CASCADE;
DROP TABLE IF EXISTS public.siem_filter_types CASCADE;
DROP TABLE IF EXISTS public.siem_external_systems CASCADE;
DROP TABLE IF EXISTS public.siem_errors CASCADE;
DROP TABLE IF EXISTS public.siem_actions CASCADE;

-- Drop signatures tables
DROP TABLE IF EXISTS public.signatures_test CASCADE;
DROP TABLE IF EXISTS public.signatures_overlay CASCADE;

-- Drop sensors settings
DROP TABLE IF EXISTS public.sensors_settings CASCADE;

-- Drop schema migrations
DROP TABLE IF EXISTS public.schema_migrations CASCADE;

-- Drop persons and related tables
DROP TABLE IF EXISTS public.persons CASCADE;

-- Drop organizations
DROP TABLE IF EXISTS public.orgs CASCADE;

-- Drop objects and related tables
DROP TABLE IF EXISTS public.objects_persons CASCADE;
DROP TABLE IF EXISTS public.objects_resources CASCADE;
DROP TABLE IF EXISTS public.objects_resource_types CASCADE;
DROP TABLE IF EXISTS public.objects_networks CASCADE;

-- Drop IOCs tables
DROP TABLE IF EXISTS public.iocs_whitelist CASCADE;
DROP TABLE IF EXISTS public.iocs_types CASCADE;
DROP TABLE IF EXISTS public.iocs_sources_statuses CASCADE;
DROP TABLE IF EXISTS public.iocs_sources CASCADE;
DROP TABLE IF EXISTS public.iocs CASCADE;

-- Drop IAM tables
DROP TABLE IF EXISTS public.iam_users_roles CASCADE;
DROP TABLE IF EXISTS public.iam_users CASCADE;
DROP TABLE IF EXISTS public.iam_roles_permissions CASCADE;
DROP TABLE IF EXISTS public.iam_roles CASCADE;
DROP TABLE IF EXISTS public.iam_permissions CASCADE;
DROP TABLE IF EXISTS public.iam_orgs CASCADE;

-- Drop GeoIP
DROP TABLE IF EXISTS public.geoip CASCADE;

-- Drop events and related tables
DROP TABLE IF EXISTS public.events CASCADE;
DROP TABLE IF EXISTS public.event_ttps CASCADE;
DROP TABLE IF EXISTS public.event_rules CASCADE;
DROP TABLE IF EXISTS public.event_sources CASCADE;
DROP TABLE IF EXISTS public.event_home_ipaddresses CASCADE;
DROP TABLE IF EXISTS public.event_ext_ipaddresses CASCADE;
DROP TABLE IF EXISTS public.event_attributes CASCADE;
DROP TABLE IF EXISTS public.event_attachments CASCADE;

-- Drop dictionary tables
DROP TABLE IF EXISTS public.dict_vulnerabilities CASCADE;
DROP TABLE IF EXISTS public.dict_softs CASCADE;
DROP TABLE IF EXISTS public.dict_soft_types CASCADE;
DROP TABLE IF EXISTS public.dict_person_roles CASCADE;
DROP TABLE IF EXISTS public.dict_person_orders CASCADE;
DROP TABLE IF EXISTS public.dict_person_modes CASCADE;
DROP TABLE IF EXISTS public.dict_misp_attribute_types CASCADE;
DROP TABLE IF EXISTS public.dict_misp_attribute_categories CASCADE;
DROP TABLE IF EXISTS public.dict_event_types CASCADE;
DROP TABLE IF EXISTS public.dict_event_statuses CASCADE;
DROP TABLE IF EXISTS public.dict_attack_protocols CASCADE;
DROP TABLE IF EXISTS public.dict_attack_methods CASCADE;
DROP TABLE IF EXISTS public.dict_attack_levels CASCADE;
DROP TABLE IF EXISTS public.dict_attack_classes CASCADE;

-- Drop grids states
DROP TABLE IF EXISTS public.grids_states CASCADE;

-- Drop all sequences
DROP SEQUENCE IF EXISTS public.users_roles_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.users_metrics_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.users_histories_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.users_actions_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.users_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.subjects_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_geo_regions_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_geo_districts_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_geo_countries_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_geo_cities_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_tags_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_tagged_item_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_tag_object_types_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_menu_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_custom_fields_objects_types_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_custom_fields_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_custom_fields_choicesets_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_extras_changelog_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_signatures_classtypes_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_sensors_hardwares_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_sensors_groups_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_recomendations_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_nccci_classes_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_mitre_techniques_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_mitre_tactics_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_local_classtypes_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_iocs_whitelist_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_iocs_types_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_iocs_sources_statuses_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_iocs_sources_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_iocs_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_detection_classtypes_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_core_content_type_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_content_type_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_assets_homenets_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_actors_subjects_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_actors_scopes_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_actors_persons_orders_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.socbox_actors_objects_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.signatures_test_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.signatures_overlay_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.siem_task_types_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.siem_task_statuses_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.siem_tasks_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.siem_filters_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.siem_filter_types_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.siem_external_systems_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.sensors_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.persons_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.orgs_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.objects_resources_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.objects_resource_types_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.objects_networks_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.objects_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.iam_users_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.iam_roles_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.iam_permissions_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.iam_orgs_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.events_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.event_ttps_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.event_rules_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.event_home_ipaddresses_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.event_ext_ipaddresses_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.event_attributes_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.event_attachments_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_vulnerabilities_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_softs_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_recomendations_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_regions_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_mitre_techniques_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_mitre_tactics_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_misp_attribute_types_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_misp_attribute_categories_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_event_statuses_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_districts_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_cities_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.dict_attack_methods_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.assets_crud_grids_id_seq CASCADE;
