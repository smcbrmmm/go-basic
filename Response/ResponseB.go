package response

type responseB struct {
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

func ReturnResponse(name, fullName string) responseB {

	return responseB{Name: name, FullName: fullName}
}
