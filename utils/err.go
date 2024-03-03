package utils

func ErrCheck(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
