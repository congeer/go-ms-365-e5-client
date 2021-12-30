package request

type CreateFolderRequest struct {
	Name   string `json:"name"`
	Folder struct {
	} `json:"folder"`
	ConflictBehavior string `json:"@microsoft.graph.conflictBehavior"`
}

func NewCreateFolderRequest(name string) CreateFolderRequest {
	return CreateFolderRequest{
		Name:             name,
		ConflictBehavior: "rename",
	}
}
