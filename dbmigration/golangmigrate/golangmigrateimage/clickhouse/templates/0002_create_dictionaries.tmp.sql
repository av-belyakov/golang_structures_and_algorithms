--- DICTIONARY ${CLICKHOUSE_DB}.siem2_iocsources
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_iocsources
(
    id UInt64,
    name String DEFAULT '',
    description String DEFAULT '',
    weigth UInt16 DEFAULT 0,
    updated_at DateTime,
    created_at DateTime,
    updated_by UInt16 DEFAULT 1,
    created_by UInt16 DEFAULT 1
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'iocs_sources'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_geoip
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_geoip
(
    inetnum_key String,                 -- ключ
    inetnum String DEFAULT '',          -- атрибут (копия ключа)
    netname String DEFAULT '',
    description String DEFAULT '',
    alpha2 String DEFAULT '',
    city String DEFAULT '',
    region String DEFAULT '',
    asn UInt32 DEFAULT 0,
    latitude Float64 DEFAULT 0.0,
    longitude Float64 DEFAULT 0.0,
    created_at DateTime,
    updated_at DateTime,
    updated_by UInt16 DEFAULT 1,
    created_by UInt16 DEFAULT 1
)
PRIMARY KEY inetnum_key
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    QUERY 'SELECT inetnum AS inetnum_key, inetnum, netname, description, alpha2, city, region, asn, latitude, longitude, created_at, updated_at, updated_by, created_by FROM geoip'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(IP_TRIE());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_objects
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_objects
(
    id UInt64,
    name String DEFAULT '',
    full_name String DEFAULT '',
    scope_id UInt16 DEFAULT 0,
    latitude Float64 DEFAULT 0,
    longitude Float64 DEFAULT 0,
    subject_id UInt32 DEFAULT 0,
    disabled_at DateTime DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_actors_objects'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_externalsystems
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_externalsystems
(
    id UInt64,
    name String DEFAULT '',
    description String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'siem_external_systems'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_priorities
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_priorities
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_detection_attacks_priorities'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_regions
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_regions
(
    id UInt64,
    name String DEFAULT '',
    district_id UInt16 DEFAULT 0,
    geocode String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_geo_regions'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_iocs
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_iocs
(
    id UInt64,
    value String DEFAULT '',
    ioc_type_id UInt16 DEFAULT 0,
    ioc_source_id UInt16 DEFAULT 0,
    active UInt8 DEFAULT 1,
    threat_level_id UInt8 DEFAULT 4,
    reference String DEFAULT '',
    score Float64 DEFAULT 0.0,
    description String DEFAULT '',
    rating_up UInt32 DEFAULT 0,
    rating_down UInt32 DEFAULT 0,
    updated_at DateTime,
    created_at DateTime,
    updated_by UInt16 DEFAULT 1,
    created_by UInt16 DEFAULT 1
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'iocs'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_eventsources
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_eventsources
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'event_sources'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_localclasstypes
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_localclasstypes
(
    id UInt64,
    name String DEFAULT '',
    description String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'public.socbox_detection_local_classtypes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_countries
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_countries
(
    alpha2 String,
    name String DEFAULT '',
    full_name String DEFAULT ''
)
PRIMARY KEY alpha2
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_geo_countries'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(COMPLEX_KEY_HASHED());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_signatures
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_signatures
(
    i_sid UInt64,
    i_gid UInt64,
    id UInt64 DEFAULT 0,
    i_priority UInt8 DEFAULT 0,
    s_msg String DEFAULT '',
    s_rule_body String DEFAULT '',
    s_source_name String DEFAULT '',
    rating_up UInt16 DEFAULT 0,
    rating_down UInt16 DEFAULT 0,
    s_ref String DEFAULT '',
    classtype_id UInt8 DEFAULT 0,
    local_classtype_id UInt8 DEFAULT 0,
    nccci_classtype_id UInt8 DEFAULT 0,
    created_at DateTime,
    updated_at DateTime,
    created_by UInt16 DEFAULT 1,
    updated_by UInt16 DEFAULT 1,
    i_b_rec UInt8 DEFAULT 1,
    i_b_com UInt8 DEFAULT 0
)
PRIMARY KEY i_sid, i_gid
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'signatures_test'
))
LIFETIME(MIN 1200 MAX 1200)
LAYOUT(COMPLEX_KEY_HASHED());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_directions
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_directions
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_detection_directions'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_homenets
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_homenets
(
    network String,
    id UInt64 DEFAULT 0,
    bandwidth UInt64 DEFAULT 0,
    object_id UInt32 DEFAULT 0,
    comment String DEFAULT '',
    created_at DateTime,
    updated_at DateTime,
    created_by UInt16 DEFAULT 1,
    updated_by UInt16 DEFAULT 1
)
PRIMARY KEY network
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_assets_homenets'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(IP_TRIE());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_resources
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_resources
(
    id UInt64,
    domain String DEFAULT '',
    url String DEFAULT '',
    object_resource_type_id UInt16 DEFAULT 0,
    description String DEFAULT '',
    object_homenets_id UInt32 DEFAULT 0,
    ip_address String DEFAULT '',
    local_ip_address String DEFAULT '',
    created_at DateTime,
    updated_at DateTime
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_assets_resources'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_classtypes
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_classtypes
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_detection_classtypes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_nccciclasses
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_nccciclasses
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'socbox_detection_nccci_classes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY ${CLICKHOUSE_DB}.siem2_scopes
CREATE DICTIONARY ${CLICKHOUSE_DB}.siem2_scopes
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST '${POSTGRES_HOST}'
    PORT ${POSTGRES_PORT}
    USER '${POSTGRES_USER}'
    PASSWORD '${POSTGRES_PASSWORD}'
    DB '${POSTGRES_DB}'
    TABLE 'public.socbox_actors_scopes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

SYSTEM RELOAD DICTIONARIES;

--------------------------------------------------------

-- TABLE ${CLICKHOUSE_DB}.alerts
CREATE TABLE ${CLICKHOUSE_DB}.alerts
(
    `object_id` UInt32,
    `start_kv_sec` UInt32,
    `start_kv_msec` UInt32,
    `end_kv_sec` UInt32,
    `end_kv_msec` UInt32,   
    `total` UInt32,
    `srcip` UInt32,
    `dstip` UInt32,
    `srcport` UInt32,
    `dstport` UInt32,
    `protocol` UInt8,
    `eth_type` UInt16,
    `gid` UInt32,
    `sid` UInt64,
    `payload` String,
    `id` UUID MATERIALIZED generateUUIDv4(),
    `adate` Date MATERIALIZED toDate(start_kv_sec),
    `idatetime` DateTime MATERIALIZED now(),
    `srccountry` String MATERIALIZED dictGetString('${CLICKHOUSE_DB}.siem2_geoip', 'alpha2', tuple(srcip)),
    `dstcountry` String MATERIALIZED dictGetString('${CLICKHOUSE_DB}.siem2_geoip', 'alpha2', tuple(dstip)),
    `srcregion` String MATERIALIZED dictGetString('${CLICKHOUSE_DB}.siem2_geoip', 'region', tuple(srcip)),
    `dstregion` String MATERIALIZED dictGetString('${CLICKHOUSE_DB}.siem2_geoip', 'region', tuple(dstip)),
    `srccity` String MATERIALIZED dictGetString('${CLICKHOUSE_DB}.siem2_geoip', 'city', tuple(srcip)),
    `dstcity` String MATERIALIZED dictGetString('${CLICKHOUSE_DB}.siem2_geoip', 'city', tuple(dstip)),
    `priority` UInt8 MATERIALIZED dictGetUInt8('${CLICKHOUSE_DB}.siem2_signatures', 'i_priority', (toUInt64(sid), toUInt64(gid))),
    `direction` UInt8 MATERIALIZED if(dictHas('${CLICKHOUSE_DB}.siem2_homenets', tuple(srcip)), if(dictHas('${CLICKHOUSE_DB}.siem2_homenets', tuple(dstip)), caseWithExpression(dictGetUInt32('${CLICKHOUSE_DB}.siem2_homenets', 'object_id', tuple(srcip)), 0, 0, if(dictGetUInt32('${CLICKHOUSE_DB}.siem2_homenets', 'object_id', tuple(srcip)) = object_id, 2, 1)), 2), caseWithExpression(dictHas('${CLICKHOUSE_DB}.siem2_homenets', tuple(dstip)), 1, 1, 0))
)
ENGINE = MergeTree
PARTITION BY adate
PRIMARY KEY (object_id, start_kv_sec, sid, srcip, dstip)
ORDER BY (object_id, start_kv_sec, sid, srcip, dstip)
SETTINGS index_granularity = 8192;

-- TABLE ${CLICKHOUSE_DB}.alerts_assigned
CREATE TABLE ${CLICKHOUSE_DB}.alerts_assigned
(
    `dt_assigned` DateTime MATERIALIZED toDateTime(now()),
    `id` UUID,
    `user_id` UInt16,
    `isAnalyzed` UInt8
)
ENGINE = MergeTree
ORDER BY dt_assigned
TTL dt_assigned + toIntervalMonth(1)
SETTINGS index_granularity = 8192;

CREATE TABLE ${CLICKHOUSE_DB}.alerts_exported
(
    `id` UUID,
    `start_kv_sec` UInt32,
    `end_kv_sec` UInt32,
    `user_id` UInt16,
    `external_system_id` UInt16,
    `external_id` String,
    `adate` Date MATERIALIZED toDate(start_kv_sec),
    `edate` Date MATERIALIZED now(),
    `edatetime` DateTime MATERIALIZED now()
)
ENGINE = ReplacingMergeTree
PRIMARY KEY (edate, id)
ORDER BY (edate, id)
SETTINGS index_granularity = 8192;

-- TABLE ${CLICKHOUSE_DB}.alerts_viewed
CREATE TABLE ${CLICKHOUSE_DB}.alerts_viewed
(
    `dt_viewed` DateTime MATERIALIZED toDateTime(now()),
    `id` UUID,
    `user_id` UInt16
)
ENGINE = MergeTree
ORDER BY dt_viewed
TTL dt_viewed + toIntervalWeek(1)
SETTINGS index_granularity = 8192;

-- TABLE ${CLICKHOUSE_DB}.classtypes
CREATE TABLE ${CLICKHOUSE_DB}.classtypes
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_classtypes);

-- TABLE ${CLICKHOUSE_DB}.directions
CREATE TABLE ${CLICKHOUSE_DB}.directions
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_directions);

-- TABLE ${CLICKHOUSE_DB}.event_rules
CREATE TABLE ${CLICKHOUSE_DB}.event_rules
(
    `id` UInt64,
    `event_id` UInt64,
    `signatures_sid_id` UInt64,
    `signatures_gid_id` UInt64,
    `alert_count` UInt64
)
ENGINE = PostgreSQL('${POSTGRES_HOST}:${POSTGRES_PORT}', '${POSTGRES_DB}', 'event_rules', '${POSTGRES_USER}', '${POSTGRES_PASSWORD}');

-- TABLE ${CLICKHOUSE_DB}.event_ttps
CREATE TABLE ${CLICKHOUSE_DB}.event_ttps
(
    `id` UInt64,
    `event_id` UInt64,
    `mitre_techniques_id` UInt16
)
ENGINE = PostgreSQL('${POSTGRES_HOST}:${POSTGRES_PORT}', '${POSTGRES_DB}', 'event_rules', '${POSTGRES_USER}', '${POSTGRES_PASSWORD}');

-- TABLE ${CLICKHOUSE_DB}.events
CREATE TABLE ${CLICKHOUSE_DB}.events
(
    `id` UInt64,
    `e_start` DateTime,
    `e_end` DateTime,
    `nccci_classtype_id` UInt8,
    `recomendation_id` UInt8
)
ENGINE = PostgreSQL('${POSTGRES_HOST}:${POSTGRES_PORT}', '${POSTGRES_DB}', 'event_rules', '${POSTGRES_USER}', '${POSTGRES_PASSWORD}');
    
-- TABLE ${CLICKHOUSE_DB}.eventsources
CREATE TABLE ${CLICKHOUSE_DB}.eventsources
(
    `id` UInt64,
    `name` String
)
ENGINE = PostgreSQL('${POSTGRES_HOST}:${POSTGRES_PORT}', '${POSTGRES_DB}', 'event_sources', '${POSTGRES_USER}', '${POSTGRES_PASSWORD}');
    
-- TABLE ${CLICKHOUSE_DB}.iocs
CREATE TABLE ${CLICKHOUSE_DB}.iocs
(
    `id` UInt64,
    `value` String,
    `ioc_type_id` UInt16,
    `ioc_source_id` UInt16,
    `active` UInt8,
    `threat_level_id` UInt8,
    `reference` String,
    `score` Float64,
    `description` String,
    `rating_up` UInt32,
    `rating_down` UInt32,
    `created_at` DateTime,
    `updated_at` DateTime
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_iocs);
    
--- TABLE ${CLICKHOUSE_DB}.iocsources
CREATE TABLE ${CLICKHOUSE_DB}.iocsources
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_iocsources);

--- TABLE ${CLICKHOUSE_DB}.geoip
CREATE TABLE ${CLICKHOUSE_DB}.geoip
(
    `inetnum` String,
    `netname` String,
    `description` String,
    `alpha2` String,
    `city` String,
    `region` String,
    `asn` UInt32,
    `latitude` Float64,
    `longitude` Float64,
    `created_at` DateTime,
    `updated_at` DateTime
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_geoip);

-- TABLE ${CLICKHOUSE_DB}.localclasstypes
CREATE TABLE ${CLICKHOUSE_DB}.localclasstypes
(
    `id` UInt64,
    `name` String,
    `description` String
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_localclasstypes);

-- TABLE ${CLICKHOUSE_DB}.nccciclasses
CREATE TABLE ${CLICKHOUSE_DB}.nccciclasses
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_nccciclasses);

--- TABLE ${CLICKHOUSE_DB}.objects
CREATE TABLE ${CLICKHOUSE_DB}.objects
(
    `id` UInt64,
    `name` String,
    `full_name` String,
    `scope_id` UInt16,
    `latitude` Float64,
    `longitude` Float64,
    `subject_id` UInt32,
    `disabled_at` DateTime
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_objects);

-- TABLE ${CLICKHOUSE_DB}.objects_homenets
CREATE TABLE ${CLICKHOUSE_DB}.objects_homenets
(
    `network` String,
    `bandwidth` UInt64,
    `created_at` DateTime,
    `updated_at` DateTime,
    `comment` String,
    `object_id` UInt32,
    `id` UInt64
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_homenets);

CREATE TABLE ${CLICKHOUSE_DB}.objects_resources
(
    `id` UInt64,
    `domain` String,
    `url` String,
    `object_resource_type_id` UInt16,
    `description` String,
    `object_homenets_id` UInt32,
    `ip_address` String,
    `local_ip_address` String,
    `created_at` DateTime,
    `updated_at` DateTime
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_resources);

-- TABLE ${CLICKHOUSE_DB}.regions
CREATE TABLE ${CLICKHOUSE_DB}.regions
(
    `id` UInt64,
    `name` String,
    `district_id` UInt16,
    `geocode` String
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_regions);

-- TABLE ${CLICKHOUSE_DB}.scopes
CREATE TABLE ${CLICKHOUSE_DB}.scopes
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_scopes);

-- TABLE ${CLICKHOUSE_DB}.signatures
CREATE TABLE ${CLICKHOUSE_DB}.signatures
(
    `i_sid` UInt64,
    `i_gid` UInt64,
    `id` UInt64,
    `i_priority` UInt8,
    `s_msg` String,
    `s_rule_body` String,
    `s_source_name` String,
    `rating_up` UInt16,
    `rating_down` UInt16,
    `s_ref` String,
    `classtype_id` UInt8,
    `local_classtype_id` UInt8,
    `nccci_classtype_id` UInt8,
    `i_b_rec` UInt8,
    `i_b_com` UInt8,
    `created_at` DateTime,
    `updated_at` DateTime
)
ENGINE = Dictionary(${CLICKHOUSE_DB}.siem2_signatures);

-- TABLE ${CLICKHOUSE_DB}.tasks
CREATE TABLE ${CLICKHOUSE_DB}.tasks
(
    `id` UUID,
    `tdate` Date MATERIALIZED toDate(now()),
    `tdatetime` DateTime MATERIALIZED now(),
    `user_id` UInt8,
    `object_id` UInt32,
    `srcip` UInt32,
    `dstip` UInt32,
    `port` Nullable(String) DEFAULT NULL,
    `starttimefilter` DateTime,
    `endtimefilter` DateTime
)
ENGINE = MergeTree
ORDER BY (tdate, id)
SETTINGS index_granularity = 8192;

-- TABLE ${CLICKHOUSE_DB}.tasks2
CREATE TABLE ${CLICKHOUSE_DB}.tasks2
(
    `id` UUID,
    `tdate` Date MATERIALIZED toDate(now()),
    `status` Enum8('Выполнено' = 0, 'Фильтруется' = 1, 'Скачивается' = 2, 'Ошибка' = 3, 'Создана' = 4, 'Подготовка' = 5, 'Подключение' = 6),
    `percent` UInt8,
    `error` String,
    `result` String
)
ENGINE = MergeTree
ORDER BY (tdate, id)
SETTINGS index_granularity = 8192;

-- TABLE ${CLICKHOUSE_DB}.users
CREATE TABLE ${CLICKHOUSE_DB}.users
(
    `id` UInt16,
    `login` String,
    `name1` String,
    `name2` String,
    `name3` String
)
ENGINE = PostgreSQL('${POSTGRES_HOST}:${POSTGRES_PORT}', '${POSTGRES_DB}', 'users', '${POSTGRES_USER}', '${POSTGRES_PASSWORD}');

-- TABLE ${CLICKHOUSE_DB}.users_histories
CREATE TABLE ${CLICKHOUSE_DB}.users_histories
(
    `id` UInt64,
    `user_id` UInt16,
    `expression` String,
    `updated_at` DateTime,
    `created_at` DateTime
)
ENGINE = PostgreSQL('${POSTGRES_HOST}:${POSTGRES_PORT}', '${POSTGRES_DB}', 'users_histories', '${POSTGRES_USER}', '${POSTGRES_PASSWORD}');

-- Below are the rollback commands

-- Example rollback commands:
DROP TABLE IF EXISTS alerts;
DROP TABLE IF EXISTS alerts_assigned;
DROP TABLE IF EXISTS alerts_exported;
DROP TABLE IF EXISTS alerts_viewed;
DROP TABLE IF EXISTS classtypes;
DROP TABLE IF EXISTS directions;
DROP TABLE IF EXISTS event_rules;
DROP TABLE IF EXISTS event_ttps;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS eventsources;
DROP TABLE IF EXISTS iocs;
DROP TABLE IF EXISTS iocsources;
DROP TABLE IF EXISTS geoip;
DROP TABLE IF EXISTS localclasstypes;
DROP TABLE IF EXISTS nccciclasses;
DROP TABLE IF EXISTS objects;
DROP TABLE IF EXISTS objects_homenets;
DROP TABLE IF EXISTS objects_resources;
DROP TABLE IF EXISTS regions;
DROP TABLE IF EXISTS scopes;
DROP TABLE IF EXISTS signatures;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS tasks2;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS users_histories;

-- Drop dictionaries
DROP DICTIONARY IF EXISTS siem2_iocsources;
DROP DICTIONARY IF EXISTS siem2_geoip;
DROP DICTIONARY IF EXISTS siem2_objects;
DROP DICTIONARY IF EXISTS siem2_externalsystems;
DROP DICTIONARY IF EXISTS siem2_priorities;
DROP DICTIONARY IF EXISTS siem2_regions;
DROP DICTIONARY IF EXISTS siem2_iocs;
DROP DICTIONARY IF EXISTS siem2_eventsources;
DROP DICTIONARY IF EXISTS siem2_localclasstypes;
DROP DICTIONARY IF EXISTS siem2_countries;
DROP DICTIONARY IF EXISTS siem2_signatures;
DROP DICTIONARY IF EXISTS siem2_directions;
DROP DICTIONARY IF EXISTS siem2_homenets;
DROP DICTIONARY IF EXISTS siem2_resources;
DROP DICTIONARY IF EXISTS siem2_classtypes;
DROP DICTIONARY IF EXISTS siem2_nccciclasses;
DROP DICTIONARY IF EXISTS siem2_scopes;
