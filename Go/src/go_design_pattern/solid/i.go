/*
接口隔离原则（ISP） 接口隔离原则指客户端不应被迫依赖于他们不使用的接口。换句话说，我们应该使用多个更小、更专业的接口，而不是单个、庞大的接口。
*/

package solid

type CloudStorage interface {
	UploadFile(file string) error
	DeleteFile(file string) error
	DownloadFile(file string) error
	GetFileMetadata(file string) (Metadata, error)
}

type Metadata struct {
	Size     int64
	Created  int64
	Modified int64
}

// 如果客户端只需要上传和下载文件，他们被迫依赖于他们不使用的方法（GetFileMetadata 和 DeleteFile）。为了遵循 ISP，我们可以重构代码：

type FileUploader interface {
	UploadFile(file string) error
}

type FileDownloader interface {
	DownloadFile(file string) error
}

type MetadataRetriever interface {
	GetFileMetadata(file string) (Metadata, error)
}

type FileDeleter interface {
	DeleteFile(file string) error
}
