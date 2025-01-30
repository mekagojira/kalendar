package server

import listevent "kalendar/feature/list_event"

func startController() {
	registerEndpoint("/v1/list-event", listevent.Handle)

	registerEndpoint("/list-event", listevent.Handle)
}
