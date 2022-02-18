package goth

/* Check error utility instead of the very verbose if error != nil {...} boilerplate. */
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}