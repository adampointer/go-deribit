package deribit

// GenerateSig creates the signature required for private endpoints
/*func (r *RPCRequest) GenerateSig(key, secret string) error {
	if len(key) == 0 || len(secret) == 0 {
		return errors.New("You must supply an access key and an access secret")
	}
	nonce := time.Now().UnixNano() / int64(time.Millisecond)
	sigString := fmt.Sprintf("_=%d&_ackey=%s&_acsec=%s&_action=%s", nonce, key, secret, r.Action)

	// Append args if present
	if len(r.Arguments) != 0 {
		var argsString string

		// We have to do this to sort by keys
		keys := make([]string, 0, len(r.Arguments))
		for key := range r.Arguments {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := r.Arguments[k]
			var s string

			switch t := v.(type) {
			case []SubscriptionEvent:
				var str = make([]string, len(t))
				for _, j := range t {
					str = append(str, string(j))
				}
				s = strings.Join(str, "")
			case []string:
				s = strings.Join(t, "")
			case bool:
				s = strconv.FormatBool(t)
			case int:
				s = strconv.FormatInt(int64(t), 10)
			case int64:
				s = strconv.FormatInt(t, 10)
			case float64:
				s = strconv.FormatFloat(t, 'f', -1, 64)
			case string:
				s = t
			default:
				// Absolutely panic here
				panic(fmt.Sprintf("Cannot generate sig string: Unable to handle arg of type %T", t))
			}
			argsString += fmt.Sprintf("&%s=%s", k, s)
		}
		sigString += argsString
	}
	hasher := sha256.New()
	hasher.Write([]byte(sigString))
	sigHash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	r.Sig = fmt.Sprintf("%s.%d.%s", key, nonce, sigHash)
	return nil
}*/