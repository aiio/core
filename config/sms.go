package config

type SMS struct {
	Provider  string
	AccessID  string
	AccessKey string
	AppID     string
	Sign      string
	Template  string
	Account   string
}

//switch config.GetEnv("SMS_PROVIDER") {
//	case Aliyun:
//		return NewAliyunClient(config.GetEnv("SMS_ACCESS_ID"), config.GetEnv("SMS_ACCESS_KEY"),
//			config.GetEnv("SMS_SIGN"), config.GetEnv("SMS_TEMPLATE"))
//	case TencentCloud:
//		return NewTencentClient(config.GetEnv("SMS_ACCESS_ID"), config.GetEnv("SMS_ACCESS_KEY"),
//			config.GetEnv("SMS_SIGN"), config.GetEnv("SMS_TEMPLATE"), config.GetEnv("SMS_APP_ID"))
//	case VolcEngine:
//		return NewVolcClient(config.GetEnv("SMS_ACCESS_ID"), config.GetEnv("SMS_ACCESS_KEY"),
//			config.GetEnv("SMS_SIGN"), config.GetEnv("SMS_TEMPLATE"), config.GetEnv("SMS_ACCOUNT"))
//	case Huyi:
//		return NewHuyiClient(config.GetEnv("SMS_ACCESS_ID"), config.GetEnv("SMS_ACCESS_KEY"), config.GetEnv("SMS_TEMPLATE"))
//	default:
//		return nil, fmt.Errorf("unsupported provider: %s", config.GetEnv("SMS_PROVIDER"))
//	}
