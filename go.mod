module k8s-scheduler-extender-example

go 1.17

replace (
	k8s.io/api => k8s.io/api v0.15.12
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.15.12
	k8s.io/apimachinery => k8s.io/apimachinery v0.15.12
	k8s.io/apiserver => k8s.io/apiserver v0.15.12
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.15.12
	k8s.io/client-go => k8s.io/client-go v0.15.12
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.15.12
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.15.12
	k8s.io/code-generator => k8s.io/code-generator v0.15.12
	k8s.io/component-base => k8s.io/component-base v0.15.12
	k8s.io/cri-api => k8s.io/cri-api v0.15.12
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.15.12
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.15.12
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.15.12
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.15.12
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.15.12
	k8s.io/kubectl => k8s.io/kubectl v0.15.12
	k8s.io/kubelet => k8s.io/kubelet v0.15.12
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.15.12
	k8s.io/metrics => k8s.io/metrics v0.15.12
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.15.12
)

// replace github.com/charstal/load-monitor => /home/charstal/workplace/project/graduation-project/load-monitor

require (
	github.com/charstal/load-monitor v1.15.12
	github.com/comail/colog v0.0.0-20160416085026-fba8e7b1f46c
	github.com/julienschmidt/httprouter v1.3.0
	k8s.io/api v0.21.6
	k8s.io/apimachinery v0.21.6
	k8s.io/apiserver v0.15.12
	k8s.io/klog/v2 v2.9.0
	k8s.io/kubernetes v1.15.12
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/go-logr/logr v0.4.0 // indirect
	github.com/gogo/protobuf v1.1.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d // indirect
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.12.1 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/spf13/pflag v1.0.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.11.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/apiextensions-apiserver v0.0.0 // indirect
	k8s.io/client-go v0.21.6 // indirect
	k8s.io/component-base v0.15.12 // indirect
	k8s.io/klog v0.3.1 // indirect
	k8s.io/metrics v0.19.0 // indirect
	k8s.io/utils v0.0.0-20190221042446-c2654d5206da // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)
