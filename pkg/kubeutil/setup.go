package kubeutil

import (
	"fmt"
	"log"

	"github.com/JosephZoeller/DDALITE/pkg/shellutil"
)

// SetUp deploys pods to the already made EC2 instances.
func SetUp(podCount int) error {

	// Create absolute file path to deployment yaml
	// // home, present := os.LookupEnv("HOME")
	// if present == false {
	// 	return errors.New("Home variable is not set. Can't find kubernetes/deployment.yaml")
	// }
	// Appy deployment
	out, err := shellutil.RunCommand("sudo kubectl --kubeconfig=/home/pi/.kube/config apply -f /home/pi/go/src/github.com/JosephZoeller/DDALITE/kubernetes/deployment.yaml", ".")

	if err != nil {
		log.Printf("-Deployment Failed-\n%s\n", out)
		return fmt.Errorf("Could not kubectl apply deployment")
	}

	log.Printf("-Successful Deployment-\n%s\n", out)

	// kubectl scale deployment.v1.apps/collider-deployment --relicas=```podCount```
	out, err = shellutil.RunCommand(fmt.Sprintf("sudo kubectl --kubeconfig=/home/pi/.kube/config scale deployment.v1.apps/collider-deployment --timeout 30s --replicas %d", podCount), ".")
	if err != nil {
		log.Printf("-Scaling Failed-\n%s\n", out)
		return fmt.Errorf("Could not scale deployment to match user request podCount=%d", podCount)
	}
	log.Printf("-Sucessful Scaling-\n%s\n", out)

	return nil
}
