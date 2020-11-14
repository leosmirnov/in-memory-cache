package api

type API struct {
}

func New() *API {
	return &API{
		buildInfo: buildInfo,
		cfg:       cfg,
		validator: validator,

		logger: staxlog.NewLogger("http.api"),
	}
}
