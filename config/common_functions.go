package config

func AddWinFolderSlashes(str string) (string) {

	f := string(str[0])
	l := string(str[len(str) - 1])

	if f != "/" { str = "/" + str }
	if l != "/" { str += "/" }

	return str
}
