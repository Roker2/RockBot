package errors

import "github.com/sirupsen/logrus"

func SendError(err error)  {
	if err != nil {
		logrus.Println(err)
	}
}
