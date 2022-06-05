package libs

func Queryapi(apiPath string) (response string) {
	request := Buildrequest(apiPath)
	info := Makerequest(request)
	return info

}
