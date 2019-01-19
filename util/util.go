package util

//Just becaouse im lazy
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
