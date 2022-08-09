package types

type ReleaseEvent struct {
	Action  string            `json:"action,omitempty"`
	Release RepositoryRelease `json:"release,omitempty"`
	Repo    Repository        `json:"repository,omitempty"`
	Sender  User              `json:"sender,omitempty"`
}

type RepositoryRelease struct {
	TagName                string         `json:"tag_name,omitempty"`
	TargetCommitish        string         `json:"target_commitish,omitempty"`
	Name                   string         `json:"name,omitempty"`
	Body                   string         `json:"body,omitempty"`
	Draft                  bool           `json:"draft,omitempty"`
	Prerelease             bool           `json:"prerelease,omitempty"`
	DiscussionCategoryName string         `json:"discussion_category_name,omitempty"`
	ID                     int64          `json:"id,omitempty"`
	URL                    string         `json:"url,omitempty"`
	HTMLURL                string         `json:"html_url,omitempty"`
	AssetsURL              string         `json:"assets_url,omitempty"`
	Assets                 []ReleaseAsset `json:"assets,omitempty"`
	UploadURL              string         `json:"upload_url,omitempty"`
	ZipballURL             string         `json:"zipball_url,omitempty"`
	TarballURL             string         `json:"tarball_url,omitempty"`
	Author                 User           `json:"author,omitempty"`
}

type ReleaseAsset struct {
	ID                 int64  `json:"id,omitempty"`
	URL                string `json:"url,omitempty"`
	Name               string `json:"name,omitempty"`
	Label              string `json:"label,omitempty"`
	State              string `json:"state,omitempty"`
	ContentType        string `json:"content_type,omitempty"`
	Size               int    `json:"size,omitempty"`
	BrowserDownloadURL string `json:"browser_download_url,omitempty"`
	Uploader           User   `json:"uploader,omitempty"`
}
