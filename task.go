package auto_render

type TaskParams struct {
	VideoDir     string
	OutputDir    string
	CameraTyp    string
	AdjustColor  bool
	Quality      string
	EnableTop    bool
	EnableButtom bool
}

type Task struct {
	id        uint32
	Name      string
	RenderAlg string
	Priority  int32
	Params    TaskParams
}
