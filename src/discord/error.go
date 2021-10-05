package discord

func IsMissingAccessError(err error) bool {
	return err.Error() == "HTTP 403 Forbidden, {\"message\": \"Missing Access\", \"code\": 50001}"
}
