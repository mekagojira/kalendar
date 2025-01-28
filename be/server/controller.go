package server

import listevent "kalendar/feature/list_event"

func startController() {
	registerEndpoint("/list-event", listevent.Handle)
}
