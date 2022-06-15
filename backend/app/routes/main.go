package routes

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	authRoutes *AuthRoutes,
) Routes {
	return Routes{
		authRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
