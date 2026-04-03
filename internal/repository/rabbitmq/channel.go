package rabbitmq

import "log"

func Createchannel() error {

	 chl,chlerr := Rabbitmq.Rabbitmqconn.Channel()

	 if chlerr != nil {
		log.Println("Channel Creation Error",chlerr)
		return chlerr 
	 }

	 defer chl.Close()

	 Createqueue(chl)

	 return nil 
}