package main

import (
	"github.com/Javlopez/kube-user/cmd"
)

func main() {
	cmd.Execute()

	// cobra.OnInitialize(initConfig)
}

func initConfig() {

	/*
		// Create a Kubernetes client configuration from the kubeconfig file.
		kubeconfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
		if err != nil {
			panic(err)
		}

		fmt.Println("Get Kubernetes pods")

		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("error getting user home dir: %v\n", err)
			os.Exit(1)
		}
		kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
		fmt.Printf("Using kubeconfig: %s\n", kubeConfigPath)

		kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		if err != nil {
			fmt.Printf("error getting Kubernetes config: %v\n", err)
			os.Exit(1)
		}

		clientset, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			fmt.Printf("error getting Kubernetes clientset: %v\n", err)
			os.Exit(1)
		}

		pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("error getting pods: %v\n", err)
			os.Exit(1)
		}
		for _, pod := range pods.Items {
			fmt.Printf("Pod name: %s\n", pod.Name)
		}*/
}

/*
kube-user --

*/
