module github.com/grafana/grafana

go 1.17

// Override xorm's outdated go-mssqldb dependency, since we can't upgrade to current xorm (due to breaking changes).
// We need a more current go-mssqldb so we get rid of a version of apache/thrift with vulnerabilities.
// Also, use our fork with fixes for unimplemented methods (required for Go 1.16).
replace github.com/denisenkom/go-mssqldb => github.com/grafana/go-mssqldb v0.0.0-20210326084033-d0ce3c521036

// Override k8s.io/client-go outdated dependency, which is an indirect dependency of grafana/loki.
// It's also present on grafana/loki's go.mod so we'll need till it gets updated.
replace k8s.io/client-go => k8s.io/client-go v0.22.1

replace github.com/russellhaering/goxmldsig@v1.1.0 => github.com/russellhaering/goxmldsig v1.1.1

require (
	cloud.google.com/go/storage v1.18.2
	cuelang.org/go v0.4.0
	github.com/Azure/azure-sdk-for-go v59.3.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.19.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.10.0
	github.com/Azure/go-autorest/autorest v0.11.22
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/semver v1.5.0
	github.com/VividCortex/mysqlerr v0.0.0-20170204212430-6c6b55f8796f
	github.com/aws/aws-sdk-go v1.42.8
	github.com/beevik/etree v1.1.0
	github.com/benbjohnson/clock v1.1.0
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/centrifugal/centrifuge v0.19.0
	github.com/cortexproject/cortex v1.10.1-0.20211104100946-3f329a21cad4
	github.com/crewjam/saml v0.4.6-0.20210521115923-29c6295245bd
	github.com/davecgh/go-spew v1.1.1
	github.com/denisenkom/go-mssqldb v0.11.0
	github.com/dop251/goja v0.0.0-20210804101310-32956a348b49
	github.com/fatih/color v1.12.0
	github.com/gchaincl/sqlhooks v1.3.0
	github.com/getsentry/sentry-go v0.10.0
	github.com/go-openapi/strfmt v0.20.2
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-stack/stack v1.8.0
	github.com/gobwas/glob v0.2.3
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/snappy v0.0.4
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/websocket v1.4.2
	github.com/gosimple/slug v1.9.0
	github.com/grafana/cuetsy v0.0.0-20211119211437-8c25464cc9bf
	github.com/grafana/grafana-aws-sdk v0.10.1
	github.com/grafana/grafana-plugin-sdk-go v0.129.0
	github.com/grafana/loki v1.6.2-0.20211105203243-3987933cdb1b
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/go-hclog v0.16.2
	github.com/hashicorp/go-plugin v1.4.3
	github.com/hashicorp/go-version v1.3.0
	github.com/influxdata/influxdb-client-go/v2 v2.6.0
	github.com/influxdata/line-protocol v0.0.0-20210311194329-9aa0e372d097
	github.com/jmespath/go-jmespath v0.4.0
	github.com/json-iterator/go v1.1.12
	github.com/jung-kurt/gofpdf v1.16.2
	github.com/laher/mergefs v0.1.1
	github.com/lib/pq v1.10.4
	github.com/linkedin/goavro/v2 v2.10.0
	github.com/m3db/prometheus_remote_client_golang v0.4.4
	github.com/magefile/mage v1.12.1
	github.com/mattn/go-isatty v0.0.16
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
	github.com/ohler55/ojg v1.12.9
	github.com/opentracing/opentracing-go v1.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/alertmanager v0.23.1-0.20211116083607-e2a10119aaf7
	github.com/prometheus/client_golang v1.12.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.32.1
	github.com/prometheus/prometheus v1.8.2-0.20211011171444-354d8d2ecfac
	github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967
	github.com/robfig/cron/v3 v3.0.1
	github.com/russellhaering/goxmldsig v1.1.1
	github.com/stretchr/testify v1.7.1
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	github.com/ua-parser/uap-go v0.0.0-20211112212520-00c877edfe0f
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli/v2 v2.3.0
	github.com/vectordotdev/go-datemath v0.1.1-0.20220323213446-f3954d0b18ae
	github.com/weaveworks/common v0.0.0-20211015155308-ebe5bdc2c89e
	github.com/xorcare/pointer v1.1.0
	github.com/yudai/gojsondiff v1.0.0
	go.opentelemetry.io/collector v0.31.0
	go.opentelemetry.io/collector/model v0.50.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	golang.org/x/crypto v0.1.0
	golang.org/x/exp v0.0.0-20210220032938-85be41e4509f
	golang.org/x/net v0.1.0
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11
	golang.org/x/tools v0.2.0
	gonum.org/v1/gonum v0.9.3
	google.golang.org/api v0.60.0
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/ldap.v3 v3.1.0
	gopkg.in/mail.v2 v2.3.1
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
	xorm.io/builder v0.3.6
	xorm.io/core v0.7.3
	xorm.io/xorm v0.8.2
)

require (
	cloud.google.com/go v0.97.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v0.7.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/FZambia/eagle v0.0.1 // indirect
	github.com/FZambia/sentinel v1.1.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/alecthomas/units v0.0.0-20210912230133-d1bdfacee922 // indirect
	github.com/andybalholm/brotli v1.0.3
	github.com/apache/arrow/go/arrow v0.0.0-20211112161151-bc219186db40 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/centrifugal/protocol v0.7.6 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/cockroachdb/apd/v2 v2.0.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/dennwc/varint v1.0.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dlclark/regexp2 v1.4.1-0.20201116162257-a2a8dda75c91 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/emicklei/proto v1.6.15 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-kit/log v0.2.0
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-openapi/analysis v0.20.1 // indirect
	github.com/go-openapi/errors v0.20.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/loads v0.20.2
	github.com/go-openapi/runtime v0.19.29 // indirect
	github.com/go-openapi/spec v0.20.4
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-openapi/validate v0.20.2 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/status v1.1.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/flatbuffers v2.0.0+incompatible // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grafana/grafana-google-sdk-go v0.0.0-20211104130251-b190293eaf58
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.1-0.20191002090509-6af20e3a5340 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	github.com/hashicorp/memberlist v0.3.1 // indirect
	github.com/hashicorp/yamux v0.0.0-20210826001029-26ff87cf9493 // indirect
	github.com/igm/sockjs-go/v3 v3.0.1 // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/mattetti/filebuffer v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/go-testing-interface v1.14.0 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/mna/redisc v1.3.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e // indirect
	github.com/opentracing-contrib/go-stdlib v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/common/sigv4 v0.1.0 // indirect
	// github.com/prometheus/exporter-toolkit v0.7.0 // indirect
	github.com/prometheus/node_exporter v1.0.0-rc.0.0.20200428091818-01054558c289 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/protocolbuffers/txtpbfmt v0.0.0-20201118171849-f6a6b3f636fc // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/rs/cors v1.8.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/segmentio/encoding v0.3.2
	github.com/sercand/kuberesolver v2.4.0+incompatible // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/weaveworks/promrus v1.2.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	go.mongodb.org/mongo-driver v1.7.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0
	go.uber.org/goleak v1.1.11-0.20210813005559-691160354723 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
)

require (
	cloud.google.com/go/kms v1.1.0
	github.com/golang-migrate/migrate/v4 v4.7.0
	gocloud.dev v0.24.0
	modernc.org/sqlite v0.0.0-00010101000000-000000000000
)

require (
	github.com/Azure/go-autorest/autorest/adal v0.9.17 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/chromedp/cdproto v0.0.0-20220208224320-6efb837e6bc2 // indirect
	github.com/cncf/udpa/go v0.0.0-20210930031921-04548b0d99d4 // indirect
	github.com/cncf/xds/go v0.0.0-20211011173535-cb28da3451f1 // indirect
	// github.com/containerd/containerd v1.5.9 // indirect
	github.com/elazarl/goproxy v0.0.0-20220115173737-adb46da277ac // indirect
	github.com/envoyproxy/go-control-plane v0.10.2-0.20220325020618-49ff273808a1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.2 // indirect
	github.com/getkin/kin-openapi v0.91.0 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/grafana/dskit v0.0.0-20211021180445-3bd016e9d7f1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.8 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/segmentio/asm v1.1.1 // indirect
	golang.org/x/mod v0.6.0 // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	modernc.org/cc/v3 v3.40.0 // indirect
	modernc.org/ccgo/v3 v3.16.13 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/strutil v1.1.3 // indirect
	modernc.org/token v1.0.1 // indirect
)

require github.com/containerd/containerd v1.5.13 // indirect

require github.com/prometheus/exporter-toolkit v0.7.3 // indirect

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/hashicorp/consul/api v1.10.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/serf v0.9.5 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	go.etcd.io/etcd/api/v3 v3.5.0 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.0 // indirect
	go.etcd.io/etcd/client/v3 v3.5.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	modernc.org/libc v1.21.4 // indirect
	modernc.org/memory v1.4.0 // indirect
)

// Use fork of crewjam/saml with fixes for some issues until changes get merged into upstream
replace github.com/crewjam/saml => github.com/grafana/saml v0.4.9-0.20220727151557-61cd9c9353fc

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.10.8

replace github.com/docker/distribution => github.com/distribution/distribution v2.8.0+incompatible

// TODO: remove once gocloud.dev releases 0.25.x
// `fileblob` implementation has buggy key ordering in 0.24.0
replace gocloud.dev v0.24.0 => github.com/google/go-cloud v0.24.1-0.20220209172924-99801bbb523a

replace (
	github.com/edsrzf/mmap-go => /home/joonl/dev/mmap-go
	github.com/hashicorp/go-sockaddr => /home/joonl/dev/go-sockaddr
	github.com/prometheus/prometheus => /home/joonl/dev/prometheus
	go.opentelemetry.io/otel/exporters/jaeger => /home/joonl/dev/jaeger
	modernc.org/libc => ./zos/libc
	modernc.org/memory => ./zos/memory
	modernc.org/sqlite => ./zos/sqlite
)
