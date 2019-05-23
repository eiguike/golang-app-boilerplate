package config

type Route struct {
  Method string
  Url string
  Handler string
}

type Routes []Route

func Get(url string, handler string) Route {
  return Route { "GET", url, handler }
}

func Delete(url string, handler string) Route {
  return Route { "DELETE", url, handler }
}

func Post(url string, handler string) Route {
  return Route { "POST", url, handler }
}

func Patch(url string, handler string) Route {
  return Route { "PATCH", url, handler }
}

func Put(url string, handler string) Route {
  return Route { "PUT", url, handler }
}
