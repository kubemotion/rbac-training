package main

import (
	"log"
	"strconv"
	"time"

	k8s "github.com/kubemotion/rbac-training/kubernetes"
)

// This binary will start and create jobs (bitcoin miner jobs as a joke) in a kubernetes cluster.
func main() {
	for prefix := 0; prefix < 10; prefix++ {
		jb := k8s.JobsSpec{
			Prefix:    strconv.Itoa(prefix),
			Namespace: "default",
			Image:     "amacneil/bitcoin",
		}

		if err := jb.Create("bitcoin-miner"); err != nil {
			log.Println(err)
		}

		time.Sleep(10 * time.Second)
	}
}
