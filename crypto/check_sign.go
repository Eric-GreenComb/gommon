package crypto

import ()

func BanerwaiApiV1CheckSign(sign, api_key string, args ...string) bool {
	if len(sign) == 0 {
		return false
	}
	total := api_key
	for _, arg := range args {
		total += arg
	}
	return CompareDoubleMd5(total, sign)
}

func BanerwaiApiV1GenSign(api_key string, args ...string) string {
	total := api_key
	for _, arg := range args {
		total += arg
	}
	return DoubleMd5(total)
}
