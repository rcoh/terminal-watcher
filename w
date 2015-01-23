wrap() {
	client -mode s -command "$*"
	$@
	client -mode e -command "$*" -status $?
}
