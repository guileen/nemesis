package modules

// Here to sort conf keys, to keep the module initial in right orders
func SortConfKeys(keys []string) []string {
	// TODO
	return keys
}

func MakeModule(key string, node Node) Module {
	switch key {
	case "json":
		return NewJSONModule(node)
	case "text":
		return NewTextModule(node)
	case "static":
		return NewStaticModule(node)
	case "proxy_pass":
		return NewProxyPassModule(node)
	default:
		panic("Unknow module: " + key)
	}
}

func MakeModules(mp Map) []Module {
	modules := make([]Module, 0, len(mp))
	keys := mp.Keys()
	keys = SortConfKeys(keys)
	for _, key := range keys {
		modules = append(modules, MakeModule(key, mp[key]))
	}
	return modules
}
