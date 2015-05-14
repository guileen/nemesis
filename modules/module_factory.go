package modules

// Here to sort conf keys, to keep the module initial in right orders
func SortConfKeys(keys []string) []string {
	// TODO
	return keys
}

func makeModule(key string, node Node) Module {
	switch key {
	case "json":
		return NewJSONModule(node)
	case "text":
		return NewTextModule(node)
	case "static":
		return NewStaticModule(node)
	case "proxy_pass":
		return NewProxyPassModule(node)
	case "routes":
		return NewRoutesModule(node)
	default:
		panic("Unknow module: " + key)
	}
}

func MakeChain(mp Map) Module {
	keys := mp.Keys()
	keys = SortConfKeys(keys)
	var firstModule Module
	var prevModule Module
	for _, key := range keys {
		module := makeModule(key, mp[key])
		if prevModule != nil {
			prevModule.SetNext(module)
		} else {
			firstModule = module
		}
		prevModule = module
	}
	return firstModule
}
