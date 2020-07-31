package methods

func PersonIsValid(r Person) ErrorMessage {

	var Error ErrorMessage

	if r.FirstName == "" {
		Error.Message = append(Error.Message, "firstName is missing")
	}

	if r.LastName == "" {
		Error.Message = append(Error.Message, "lastname is missing")
	}

	return Error
}
