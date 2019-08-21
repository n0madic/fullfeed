package fullfeed

// ExtractMethod for full text
type ExtractMethod string

var (
	// QueryMethod with goquery
	QueryMethod ExtractMethod = "query"
	// ReadabilityMethod by default
	ReadabilityMethod ExtractMethod = "readability"
	// XPathMethod with XML Path Language
	XPathMethod ExtractMethod = "xpath"
)

// Config for feed
type Config struct {
	// Base URL for all relative URLs
	// Must be specified if different from the feed domain
	BaseHref string `json:"base_href" yaml:"base_href"`

	// Feed description
	Description string `json:"description" yaml:"description"`

	// Feed cleaning filters
	Filters struct {
		// Skip article with the following words in the description
		Descriptions []string `json:"descriptions" yaml:"descriptions"`

		// Remove the following selectors from content
		Selectors []string `json:"selectors" yaml:"selectors"`

		// Remove blocks of text that contain the following words
		Text []string `json:"text" yaml:"text"`

		// Skip article with the following words in the title
		Titles []string `json:"titles" yaml:"titles"`
	} `json:"filters" yaml:"filters"`

	// Maximum number of processing workers (default 10)
	MaxWorkers uint `json:"max_workers" yaml:"max_workers"`

	// Full text extract method
	// Supported Methods: query (like jquery), xpath, readability (default)
	Method ExtractMethod `json:"method" yaml:"method"`

	// Full text extract request
	MethodRequest string `json:"method_request" yaml:"method_request"`

	// Link to the original feed
	URL string `json:"url" yaml:"url"`
}
