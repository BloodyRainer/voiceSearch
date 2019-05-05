package server

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"encoding/json"
	"github.com/BloodyRainer/voiceSearch/df"
	"github.com/BloodyRainer/voiceSearch/intentHandlers"
)

type dfReqHandler struct{}

func (rcv *dfReqHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var dfReq *df.Request
	var dfRes *df.Response
	var err error

	ctx := appengine.NewContext(r)

	// check and log dialogflow-post-request
	if r.Method == http.MethodPost {
		body, err := readPostBody(r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		logPostBody(ctx, body)

		dfReq, err = df.MakeDfRequest(body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	// dispatch intents in order to assemble dialogflow-response
	intent := dfReq.QueryResult.Intent.DisplayName
	log.Infof(ctx, "intent-name is: "+intent)

	switch intent {
	case "search_product_intend":
		dfRes, err = intentHandlers.SearchProduct(dfReq)
		if err != nil {
			log.Errorf(ctx, "failed to answer to name of player one: "+err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	default:
		log.Errorf(ctx, "unknown intent: "+intent)
		return
	}

	// respond to dialogflow
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dfRes)
}