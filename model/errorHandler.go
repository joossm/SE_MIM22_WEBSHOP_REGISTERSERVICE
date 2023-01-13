package model

func errorHandler(err error) {
	if err != nil {
		print(err)
	}
}
