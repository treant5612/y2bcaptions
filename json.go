package y2bcaptions

type ResponseContext struct {
	Captions struct {
		PlayerCaptionsRenderer struct {
			BaseUrl string `json:"baseUrl"`
		}
		PlayerCaptionsTracklistRenderer struct {
			CaptionTracks []struct {
				BaseUrl        string `json:"baseUrl"`
				IsTranslatable bool   `json:"isTranslatable"`
				LanguageCode   string `json:"languageCode"`
				Kind           string `json:"kind"`
				Name           struct {
					SimpleText string `json:"simpleText"`
				} `json:"name"`
				VssId string `json:"vssId"`
			} `json:"captionTracks"`
		} `json:"playerCaptionsTracklistRenderer"`
	} `json:"captions"`
}
