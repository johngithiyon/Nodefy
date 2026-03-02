package docker

import (
	"github.com/moby/moby/client"
)

var Dockerclient *client.Client

func DockerConn() error {

	    dockerclient,dockerclienterr := client.New(
			client.FromEnv,
		)

		if dockerclienterr != nil {
			 return dockerclienterr
		}

		Dockerclient = dockerclient

		return nil 
	   
}