--- DICTIONARY siem.siem2_iocsources
CREATE DICTIONARY siem.siem2_iocsources
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
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'iocs_sources'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY siem.siem2_geoip
CREATE DICTIONARY siem.siem2_geoip
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
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    QUERY 'SELECT inetnum AS inetnum_key, inetnum, netname, description, alpha2, city, region, asn, latitude, longitude, created_at, updated_at, updated_by, created_by FROM geoip'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(IP_TRIE());

--- DICTIONARY siem.siem2_objects
CREATE DICTIONARY siem.siem2_objects
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
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_actors_objects'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY siem.siem2_externalsystems
CREATE DICTIONARY siem.siem2_externalsystems
(
    id UInt64,
    name String DEFAULT '',
    description String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'siem_external_systems'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_priorities
CREATE DICTIONARY siem.siem2_priorities
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_detection_attacks_priorities'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_regions
CREATE DICTIONARY siem.siem2_regions
(
    id UInt64,
    name String DEFAULT '',
    district_id UInt16 DEFAULT 0,
    geocode String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_geo_regions'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_iocs
CREATE DICTIONARY siem.siem2_iocs
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
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'iocs'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_eventsources
CREATE DICTIONARY siem.siem2_eventsources
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'event_sources'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY siem.siem2_localclasstypes
CREATE DICTIONARY siem.siem2_localclasstypes
(
    id UInt64,
    name String DEFAULT '',
    description String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'public.socbox_detection_local_classtypes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_countries
CREATE DICTIONARY siem.siem2_countries
(
    alpha2 String,
    name String DEFAULT '',
    full_name String DEFAULT ''
)
PRIMARY KEY alpha2
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_geo_countries'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(COMPLEX_KEY_HASHED());

--- DICTIONARY siem.siem2_signatures
CREATE DICTIONARY siem.siem2_signatures
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
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'signatures_test'
))
LIFETIME(MIN 1200 MAX 1200)
LAYOUT(COMPLEX_KEY_HASHED());

--- DICTIONARY siem.siem2_directions
CREATE DICTIONARY siem.siem2_directions
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_detection_directions'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_homenets
CREATE DICTIONARY siem.siem2_homenets
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
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_assets_homenets'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(IP_TRIE());

--- DICTIONARY siem.siem2_resources
CREATE DICTIONARY siem.siem2_resources
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
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_assets_resources'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(HASHED());

--- DICTIONARY siem.siem2_classtypes
CREATE DICTIONARY siem.siem2_classtypes
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_detection_classtypes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_nccciclasses
CREATE DICTIONARY siem.siem2_nccciclasses
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'socbox_detection_nccci_classes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

--- DICTIONARY siem.siem2_scopes
CREATE DICTIONARY siem.siem2_scopes
(
    id UInt64,
    name String DEFAULT ''
)
PRIMARY KEY id
SOURCE(POSTGRESQL(
    HOST 'postgres'
    PORT 5432
    USER 'soc'
    PASSWORD 'yTa89-nMsf-GAbs2'
    DB 'siem'
    TABLE 'public.socbox_actors_scopes'
))
LIFETIME(MIN 600 MAX 600)
LAYOUT(FLAT());

SYSTEM RELOAD DICTIONARIES;

--------------------------------------------------------

-- TABLE siem.alerts
CREATE TABLE siem.alerts
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
    `srccountry` String MATERIALIZED dictGetString('siem.siem2_geoip', 'alpha2', tuple(srcip)),
    `dstcountry` String MATERIALIZED dictGetString('siem.siem2_geoip', 'alpha2', tuple(dstip)),
    `srcregion` String MATERIALIZED dictGetString('siem.siem2_geoip', 'region', tuple(srcip)),
    `dstregion` String MATERIALIZED dictGetString('siem.siem2_geoip', 'region', tuple(dstip)),
    `srccity` String MATERIALIZED dictGetString('siem.siem2_geoip', 'city', tuple(srcip)),
    `dstcity` String MATERIALIZED dictGetString('siem.siem2_geoip', 'city', tuple(dstip)),
    `priority` UInt8 MATERIALIZED dictGetUInt8('siem.siem2_signatures', 'i_priority', (toUInt64(sid), toUInt64(gid))),
    `direction` UInt8 MATERIALIZED if(dictHas('siem.siem2_homenets', tuple(srcip)), if(dictHas('siem.siem2_homenets', tuple(dstip)), caseWithExpression(dictGetUInt32('siem.siem2_homenets', 'object_id', tuple(srcip)), 0, 0, if(dictGetUInt32('siem.siem2_homenets', 'object_id', tuple(srcip)) = object_id, 2, 1)), 2), caseWithExpression(dictHas('siem.siem2_homenets', tuple(dstip)), 1, 1, 0))
)
ENGINE = MergeTree
PARTITION BY adate
PRIMARY KEY (object_id, start_kv_sec, sid, srcip, dstip)
ORDER BY (object_id, start_kv_sec, sid, srcip, dstip)
SETTINGS index_granularity = 8192;

-- TABLE siem.alerts_assigned
CREATE TABLE siem.alerts_assigned
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

CREATE TABLE siem.alerts_exported
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

-- TABLE siem.alerts_viewed
CREATE TABLE siem.alerts_viewed
(
    `dt_viewed` DateTime MATERIALIZED toDateTime(now()),
    `id` UUID,
    `user_id` UInt16
)
ENGINE = MergeTree
ORDER BY dt_viewed
TTL dt_viewed + toIntervalWeek(1)
SETTINGS index_granularity = 8192;

-- TABLE siem.classtypes
CREATE TABLE siem.classtypes
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(siem.siem2_classtypes);

-- TABLE siem.directions
CREATE TABLE siem.directions
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(siem.siem2_directions);

-- TABLE siem.event_rules
CREATE TABLE siem.event_rules
(
    `id` UInt64,
    `event_id` UInt64,
    `signatures_sid_id` UInt64,
    `signatures_gid_id` UInt64,
    `alert_count` UInt64
)
ENGINE = PostgreSQL('postgres:5432', 'siem', 'event_rules', 'soc', 'yTa89-nMsf-GAbs2');

-- TABLE siem.event_ttps
CREATE TABLE siem.event_ttps
(
    `id` UInt64,
    `event_id` UInt64,
    `mitre_techniques_id` UInt16
)
ENGINE = PostgreSQL('postgres:5432', 'siem', 'event_rules', 'soc', 'yTa89-nMsf-GAbs2');

-- TABLE siem.events
CREATE TABLE siem.events
(
    `id` UInt64,
    `e_start` DateTime,
    `e_end` DateTime,
    `nccci_classtype_id` UInt8,
    `recomendation_id` UInt8
)
ENGINE = PostgreSQL('postgres:5432', 'siem', 'event_rules', 'soc', 'yTa89-nMsf-GAbs2');
    
-- TABLE siem.eventsources
CREATE TABLE siem.eventsources
(
    `id` UInt64,
    `name` String
)
ENGINE = PostgreSQL('postgres:5432', 'siem', 'event_sources', 'soc', 'yTa89-nMsf-GAbs2');
    
-- TABLE siem.iocs
CREATE TABLE siem.iocs
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
ENGINE = Dictionary(siem.siem2_iocs);
    
--- TABLE siem.iocsources
CREATE TABLE siem.iocsources
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(siem.siem2_iocsources);

--- TABLE siem.geoip
CREATE TABLE siem.geoip
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
ENGINE = Dictionary(siem.siem2_geoip);

-- TABLE siem.localclasstypes
CREATE TABLE siem.localclasstypes
(
    `id` UInt64,
    `name` String,
    `description` String
)
ENGINE = Dictionary(siem.siem2_localclasstypes);

-- TABLE siem.nccciclasses
CREATE TABLE siem.nccciclasses
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(siem.siem2_nccciclasses);

--- TABLE siem.objects
CREATE TABLE siem.objects
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
ENGINE = Dictionary(siem.siem2_objects);

-- TABLE siem.objects_homenets
CREATE TABLE siem.objects_homenets
(
    `network` String,
    `bandwidth` UInt64,
    `created_at` DateTime,
    `updated_at` DateTime,
    `comment` String,
    `object_id` UInt32,
    `id` UInt64
)
ENGINE = Dictionary(siem.siem2_homenets);

CREATE TABLE siem.objects_resources
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
ENGINE = Dictionary(siem.siem2_resources);

-- TABLE siem.regions
CREATE TABLE siem.regions
(
    `id` UInt64,
    `name` String,
    `district_id` UInt16,
    `geocode` String
)
ENGINE = Dictionary(siem.siem2_regions);

-- TABLE siem.scopes
CREATE TABLE siem.scopes
(
    `id` UInt64,
    `name` String
)
ENGINE = Dictionary(siem.siem2_scopes);

-- TABLE siem.signatures
CREATE TABLE siem.signatures
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
ENGINE = Dictionary(siem.siem2_signatures);

-- TABLE siem.tasks
CREATE TABLE siem.tasks
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

-- TABLE siem.tasks2
CREATE TABLE siem.tasks2
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

-- TABLE siem.users
CREATE TABLE siem.users
(
    `id` UInt16,
    `login` String,
    `name1` String,
    `name2` String,
    `name3` String
)
ENGINE = PostgreSQL('postgres:5432', 'siem', 'users', 'soc', 'yTa89-nMsf-GAbs2');

-- TABLE siem.users_histories
CREATE TABLE siem.users_histories
(
    `id` UInt64,
    `user_id` UInt16,
    `expression` String,
    `updated_at` DateTime,
    `created_at` DateTime
)
ENGINE = PostgreSQL('postgres:5432', 'siem', 'users_histories', 'soc', 'yTa89-nMsf-GAbs2');

