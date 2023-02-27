package exceptions

import "log"

type DontReport struct {
}

type DontFlash struct {
}

func NotFoundHttpException(err error) {

}

func RouteNotFoundException(err error) {

}

func BadRequestException(err error) {
	if err != nil {
		log.Println("An error occurred:", err)
	}
}

func ConnectionRefused(err error) {
	if err != nil {
		log.Println("Connection Refused:", err)
	}
}
