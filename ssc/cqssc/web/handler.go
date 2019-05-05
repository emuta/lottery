package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "lottery/ssc/cqssc/proto"
)

type Handler struct {
	client pb.CqsscServiceClient
}

func NewHandler(addr string) *Handler {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatalf("Failed to connect GRPC server %s", addr)
	}

	handler := Handler{
		client: pb.NewCqsscServiceClient(conn),
	}
	return &handler
}

func (h *Handler) Params(r *http.Request) map[string]string {
	params := make(map[string]string)
	values := r.URL.Query()
	for k, _ := range values {
		if v := values.Get(k); v != "" {
			params[k] = strings.TrimSpace(v)
		}
	}
	return params
}

func (h *Handler) Error(w http.ResponseWriter, body error, statusCode int) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Fprintln(w, err)
	}
}

func (h *Handler) WriteJson(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.WithError(err).Error("Failed to encode response")
		log.Errorf("%#v \n", err)
		h.Error(w, err, 500)
	}
}

func (h *Handler) GetConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	resp, err := h.client.GetConfig(r.Context(), &pb.GetConfigReq{})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetUnit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetUnit(r.Context(), &pb.GetUnitReq{Id: idInt64})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) ListUnit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req pb.FindUnitReq

	ps := h.Params(r)
	if id, ok := ps["id"]; ok {
		v, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.Id = v
	}

	if name, ok := ps["name"]; ok {
		req.Name = name
	}

	if value, ok := ps["value"]; ok {
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.Value = v
	}

	resp, err := h.client.FindUnit(r.Context(), &req)
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetCatg(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetCatg(r.Context(), &pb.GetCatgReq{Id: int32(idInt64)})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) ListCatg(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req pb.FindCatgReq

	ps := h.Params(r)

	if id, ok := ps["id"]; ok {
		v, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.Id = int32(v)
	}

	if name, ok := ps["name"]; ok {
		req.Name = name
	}

	if tag, ok := ps["tag"]; ok {
		req.Tag = tag
	}

	if pref, ok := ps["pref"]; ok {
		req.Pref = pref
	}

	resp, err := h.client.FindCatg(r.Context(), &req)
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetGroup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetGroup(r.Context(), &pb.GetGroupReq{Id: int32(idInt64)})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) ListGroup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	defer r.Body.Close()

	var req pb.FindGroupReq

	ps := h.Params(r)

	if id, ok := ps["id"]; ok {
		v, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.Id = int32(v)
	}

	if name, ok := ps["name"]; ok {
		req.Name = name
	}

	if tag, ok := ps["tag"]; ok {
		req.Tag = tag
	}

	resp, err := h.client.FindGroup(r.Context(), &req)
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetPlay(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetPlay(r.Context(), &pb.GetPlayReq{Id: int32(idInt64)})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) ListPlay(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req pb.FindPlayReq

	ps := h.Params(r)

	if id, ok := ps["id"]; ok {
		v, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.Id = int32(v)
	}

	if name, ok := ps["name"]; ok {
		req.Name = name
	}

	if tag, ok := ps["tag"]; ok {
		req.Tag = tag
	}

	if pref, ok := ps["pref"]; ok {
		req.Pref = pref
	}

	if active, ok := ps["active"]; ok {
		req.Active = active
	}

	if pr, ok := ps["pr"]; ok {
		v, err := strconv.ParseInt(pr, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.Pr = int32(v)
	}

	if catgId, ok := ps["catg_id"]; ok {
		v, err := strconv.ParseInt(catgId, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.CatgId = int32(v)
	}

	if groupId, ok := ps["group_id"]; ok {
		v, err := strconv.ParseInt(groupId, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.GroupId = int32(v)
	}

	if unitId, ok := ps["unit_id"]; ok {
		v, err := strconv.ParseInt(unitId, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
			return
		}
		req.UnitId = v
	}

	resp, err := h.client.FindPlay(r.Context(), &req)
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetTerm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetTerm(r.Context(), &pb.GetTermReq{Id: idInt64})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) ListTerm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	/*
		int64 id = 1;
	    string code = 2;
	    google.protobuf.Timestamp start_from = 3;
	    google.protobuf.Timestamp end_to = 4;
	    int32 limit = 5;
	    int32 offset = 6;
	    string order_by = 7;
	*/
	ps := h.Params(r)
	var req pb.FindTermReq

	if id, ok := ps["id"]; ok {
		v, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
		}
		req.Id = v
	}

	if code, ok := ps["code"]; ok {
		req.Code = code
	}

	if startFrom, ok := ps["start_from"]; ok {
		t, err := time.Parse(TimeFormat, startFrom)
		if err != nil {
			if err != nil {
				log.Error(err)
				h.Error(w, err, 400)
			}
		}

		pbt, err := ptypes.TimestampProto(t)
		if err != nil {
			if err != nil {
				log.Error(err)
				h.Error(w, err, 400)
			}
		}
		req.StartFrom = pbt
	}

	if endTo, ok := ps["end_to"]; ok {
		t, err := time.Parse(TimeFormat, endTo)
		if err != nil {
			if err != nil {
				log.Error(err)
				h.Error(w, err, 400)
			}
		}
		pbt, err := ptypes.TimestampProto(t)
		if err != nil {
			if err != nil {
				log.Error(err)
				h.Error(w, err, 400)
			}
		}
		req.EndTo = pbt
	}

	if limit, ok := ps["limit"]; ok {
		v, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
		}
		req.Limit = int32(v)
	}

	if offset, ok := ps["offset"]; ok {
		v, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			log.Error(err)
			h.Error(w, err, 400)
		}
		req.Offset = int32(v)
	}

	if orderBy, ok := ps["order_by"]; ok {
		req.OrderBy = orderBy
	}

	resp, err := h.client.FindTerm(r.Context(), &req)
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) CreateBet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 1, validate date first
	// 2, call grpc method
	// 3, write response
}

func (h *Handler) GetBet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetBet(r.Context(), &pb.GetBetReq{Id: idInt64})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetBetStats(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetBetStats(r.Context(), &pb.GetBetStatsReq{Id: idInt64})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetBetAllPlans(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	resp, err := h.client.GetBetPlan(r.Context(), &pb.GetBetPlanReq{Id: idInt64})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetBetPlan(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	planIdInt64, err := strconv.ParseInt(vars["plan_id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	log.Info(planIdInt64)

	resp, err := h.client.GetBetPlan(r.Context(), &pb.GetBetPlanReq{Id: idInt64})
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) GetBetPlanStats(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	idInt64, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	planIdInt64, err := strconv.ParseInt(vars["plan_id"], 10, 64)
	if err != nil {
		log.Error(err)
		h.Error(w, err, 400)
		return
	}

	req := pb.GetBetPlanStatsReq{Id: idInt64, PlanId: planIdInt64}
	resp, err := h.client.GetBetPlanStats(r.Context(), &req)
	if err != nil {
		h.Error(w, err, 500)
		return
	}
	h.WriteJson(w, resp)
}

func (h *Handler) ListBet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}
