package utility

//CompleteKey takes authkey and returns the key with "Bot " header
func CompleteKey(x string) string {
	return "Bot " + x
}
