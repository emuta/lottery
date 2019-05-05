package web

import (
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func (h *Handler) GetRoutes() []Route {
	routes := []Route{
		Route{
			Path:    "/config",
			Handler: h.GetConfig,
			Method:  "get",
		},
		Route{
			Path:    "/units",
			Handler: h.ListUnit,
			Method:  "get",
		},
		Route{
			Path:    "/units/{id}",
			Handler: h.GetUnit,
			Method:  "get",
		},
		Route{
			Path:    "/catgs",
			Handler: h.ListCatg,
			Method:  "get",
		},
		Route{
			Path:    "/catgs/{id}",
			Handler: h.GetCatg,
			Method:  "get",
		},
		Route{
			Path:    "/groups",
			Handler: h.ListGroup,
			Method:  "get",
		},
		Route{
			Path:    "/groups/{id}",
			Handler: h.GetGroup,
			Method:  "get",
		},
		Route{
			Path:    "/plays",
			Handler: h.GetPlay,
			Method:  "get",
		},
		Route{
			Path:    "/plays/{id}",
			Handler: h.ListPlay,
			Method:  "get",
		},
		Route{
			Path:    "/terms",
			Handler: h.ListTerm,
			Method:  "get",
		},
		Route{
			Path:    "/terms/{id}",
			Handler: h.GetTerm,
			Method:  "get",
		},
		Route{
			Path:    "/bets",
			Handler: h.ListBet,
			Method:  "get",
		},
		Route{
			Path:    "/bets",
			Handler: h.CreateBet,
			Method:  "post",
		},
		Route{
			Path:    "/bets/{id}",
			Handler: h.GetBet,
			Method:  "get",
		},
		Route{
			Path:    "/bets/{id}/stats",
			Handler: h.GetBetStats,
			Method:  "get",
		},
		Route{
			Path:    "/bets/{id}/plans",
			Handler: h.GetBetAllPlans,
			Method:  "get",
		},
		Route{
			Path:    "/bets/{id}/plans/{plan_id}",
			Handler: h.GetBetPlan,
			Method:  "get",
		},
		Route{
			Path:    "/bets/{id}/plans/{plan_id}/stats",
			Handler: h.GetBetPlanStats,
			Method:  "get",
		},
	}

	return routes
}
