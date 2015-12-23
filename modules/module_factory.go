package modules

import "log"

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
	case "seturl":
		return NewSetUrlModule(node.(Scalar))
	default:
		panic("Unknow module: " + key)
	}
}

func MakeChain(mp Map) Module {
	log.Println("MakeChain", mp)
	keys := mp.Keys()
	keys = SortConfKeys(keys)
	var firstModule Module
	var prevModule Module
	for _, key := range keys {
		module := makeModule(key, mp[key])
		module.SetName(key)
		if prevModule != nil {
			prevModule.SetNext(module)
		} else {
			firstModule = module
		}
		prevModule = module
	}
	m := firstModule
	s := ""
	for m != nil {
		s = s + "->" + m.GetName()
		m = m.Next()
	}
	log.Println(s)
	return firstModule
}
