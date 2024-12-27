package main

import "fmt"

type RouteConfig struct {
	RouteIdentifiers   []RouteIdentifierHandler
	RouteCheckHandlers map[int]RouteCheckHandler
	DefaultRoute       int
}

type RouteIdentifierHandler func(req interface{}) (channelId int, err error)
type RouteCheckHandler func(req interface{}) error

var CNRouteConfig = RouteConfig{
	DefaultRoute: 1234,
}

func main() {
	routeConfig := &CNRouteConfig
	routeIdentifiers := routeConfig.RouteIdentifiers
	defaultRoute := routeConfig.DefaultRoute
	finalRoute := defaultRoute
	for _, identifier := range routeIdentifiers {
		route, err := identifier(nil)
		if err != nil {
			fmt.Println("route_not_found, err:", err)
		}
		if route != 0 {
			finalRoute = route
			break
		}
	}
	fmt.Println("finalRoute:", finalRoute)
}
