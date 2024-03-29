package option

type ListMessageMoveTasks struct {
	Default
	// The maximum number of results to include in the response. The default is 1,
	// which provides the most recent message movement task. The upper limit is 10.
	MaxResults int32
}

func NewListMessageMoveTasks() *ListMessageMoveTasks {
	return &ListMessageMoveTasks{}
}

func (l *ListMessageMoveTasks) SetDebugMode(b bool) *ListMessageMoveTasks {
	l.DebugMode = b
	return l
}

func (l *ListMessageMoveTasks) SetHttpClient(opt HttpClient) *ListMessageMoveTasks {
	l.HttpClient = &opt
	return l
}

func (l *ListMessageMoveTasks) SetMaxResults(i int32) *ListMessageMoveTasks {
	l.MaxResults = i
	return l
}

func GetListMessageMoveTaskByParams(opts []*ListMessageMoveTasks) *ListMessageMoveTasks {
	var result ListMessageMoveTasks
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		fillDefaultFields(opt.Default, &result.Default)
		if opt.MaxResults > 0 {
			result.MaxResults = opt.MaxResults
		}
	}
	if result.MaxResults == 0 {
		result.MaxResults = 1
	}
	return &result
}
