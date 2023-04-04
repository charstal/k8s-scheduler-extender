package main

import (
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/charstal/load-monitor/pkg/metricstype"
	"github.com/comail/colog"
	"github.com/julienschmidt/httprouter"

	v1 "k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

const (
	versionPath = "/version"
	apiPrefix   = "/scheduler"
	// bindPath         = apiPrefix + "/bind"
	// preemptionPath   = apiPrefix + "/preemption"
	// predicatesPrefix = apiPrefix + "/predicates"
	prioritiesPrefix = apiPrefix + "/priorities"
)

const (
	safeVarianceMargin      = 1.0
	safeVarianceSensitivity = 1.0
)

var (
	version    string // injected via ldflags at build time
	collection Collector

	collector, _ = NewCollector()

	// TruePredicate = Predicate{
	// 	Name: "always_true",
	// 	Func: func(pod v1.Pod, node v1.Node) (bool, error) {
	// 		return true, nil
	// 	},
	// }

	ZeroPriority = Prioritize{
		Name: "lvb",
		Func: func(pod v1.Pod, nodes []v1.Node) (*schedulerapi.HostPriorityList, error) {
			var priorityList schedulerapi.HostPriorityList

			priorityList = make([]schedulerapi.HostPriority, len(nodes))
			for i, node := range nodes {
				metrics, _, _ := collector.GetNodeMetrics(node.Name)
				priorityList[i] = schedulerapi.HostPriority{
					Host:  node.Name,
					Score: RequestedBasedLoadVariation(&node, &pod, metrics),
				}
			}
			return &priorityList, nil
		},
	}

	// NoBind = Bind{
	// 	Func: func(podName string, podNamespace string, podUID types.UID, node string) error {
	// 		return fmt.Errorf("This extender doesn't support Bind.  Please make 'BindVerb' be empty in your ExtenderConfig.")
	// 	},
	// }

	// EchoPreemption = Preemption{
	// 	Func: func(
	// 		_ v1.Pod,
	// 		_ map[string]*schedulerapi.Victims,
	// 		nodeNameToMetaVictims map[string]*schedulerapi.MetaVictims,
	// 	) map[string]*schedulerapi.MetaVictims {
	// 		return nodeNameToMetaVictims
	// 	},
	// }
)

// base requested
func RequestedBasedLoadVariation(
	node *v1.Node,
	pod *v1.Pod,
	nodeMetrics *[]metricstype.Metric) int {
	// calculate CPU score
	nodeName := node.Name
	score := 0

	podRequest := GetResourceRequested(pod)

	scoreFunc := func(resourceType v1.ResourceName) (float64, bool) {
		resourceStats, resourceOk := CreateResourceStats(
			*nodeMetrics, node, podRequest, resourceType,
			ResourceType2MetricTypeMap[resourceType])
		if !resourceOk {
			return 0, false
		}
		resourceScore := ComputeScore(resourceStats, safeVarianceMargin, safeVarianceSensitivity)
		log.Printf("INFO: requestedBasedLoadVariation Calculating pod %v, nodeName %s, score %f", pod, nodeName, resourceScore)
		// klog.V(6).InfoS("requestedBasedLoadVariation Calculating", "pod", klog.KObj(pod), "nodeName", nodeName,
		// 	resourceType, "Score", resourceScore)
		return resourceScore, true
	}
	// calculate total score
	totalScore := 100.0
	hasScore := false
	for tt := range ResourceType2MetricTypeMap {
		s, v := scoreFunc(tt)
		if v {
			totalScore = math.Min(totalScore, s)
			hasScore = true
		}
	}
	score = int(math.Round(totalScore))
	if !hasScore {
		score = 0
	}
	log.Printf("requestedBasedLoadVariation Calculating totalScore, pod %v, nodeName %s, total score %d", pod, node, score)
	// klog.V(6).InfoS("requestedBasedLoadVariation Calculating totalScore", "pod", klog.KObj(pod), "nodeName",
	// 	nodeName, "totalScore", score)

	return score
}

func StringToLevel(levelStr string) colog.Level {
	switch level := strings.ToUpper(levelStr); level {
	case "TRACE":
		return colog.LTrace
	case "DEBUG":
		return colog.LDebug
	case "INFO":
		return colog.LInfo
	case "WARNING":
		return colog.LWarning
	case "ERROR":
		return colog.LError
	case "ALERT":
		return colog.LAlert
	default:
		log.Printf("warning: LOG_LEVEL=\"%s\" is empty or invalid, fallling back to \"INFO\".\n", level)
		return colog.LInfo
	}
}

func main() {
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	level := StringToLevel(os.Getenv("LOG_LEVEL"))
	log.Print("Log level was set to ", strings.ToUpper(level.String()))
	colog.SetMinLevel(level)

	router := httprouter.New()
	AddVersion(router)

	// predicates := []Predicate{TruePredicate}
	// for _, p := range predicates {
	// 	AddPredicate(router, p)
	// }

	priorities := []Prioritize{ZeroPriority}
	for _, p := range priorities {
		AddPrioritize(router, p)
	}

	// AddBind(router, NoBind)

	log.Print("info: server starting on the port :80")
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}
