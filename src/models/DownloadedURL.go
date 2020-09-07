package models

import (
	"github.com/jinzhu/gorm"
	"github.com/purely-imaginary/referee-go/src/controllers"
)

// DownloadedURL ..
type DownloadedURL struct {
	URL     string `db:"url"`
	MatchID int64  `db:"match_id"`
}

// TableName .
func (u *DownloadedURL) TableName() string {
	return "downloaded_url"
}

// DoesExistsInDB .
func (u *DownloadedURL) DoesExistsInDB() int64 {
	url := &DownloadedURL{}
	err := controllers.DBEngine.First(url, "url = ?", u.URL)

	if gorm.IsRecordNotFoundError(err.Error) {
		return 0
	}

	return url.MatchID
}

// InsertIntoDB .
func (u *DownloadedURL) InsertIntoDB() {
	controllers.DBEngine.Save(u)
}
