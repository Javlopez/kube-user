package core

import (
	"fmt"
	"log"

	"github.com/Javlopez/kube-user/pkg/crypt"
	"github.com/Javlopez/kube-user/pkg/models"
	"github.com/spf13/cobra"
)

type Service struct {
	Command *cobra.Command
}

func New(command *cobra.Command) *Service {
	return &Service{
		Command: command,
	}
}

func (s Service) Run() {
	user, _ := s.Command.Flags().GetString("user")

	opts := models.Options{
		User: user,
	}

	fmt.Println("Options SET")
	fmt.Printf("opts: %+v\n", opts)

	key := crypt.New()
	err := key.Build(opts)
	if err != nil {
		log.Fatal(err)
	}

}
