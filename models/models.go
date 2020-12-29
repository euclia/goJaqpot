package models

import (
	"net/http"
	"time"
)

// ClientProperties structure
type ClientProperties struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Client structure
type Client struct {
	C ClientProperties
}

// Algorithm structure
type Algorithm struct {
	ID                 string            `json:"id,omitempty"`
	Meta               MetaInfo          `json:"meta,omitempty"`
	OntologicalClasses []string          `json:"ontologicalClasses,omitempty"`
	Visible            bool              `json:"visible,omitempty"`
	Temporary          bool              `json:"temporary,omitempty"`
	Featured           bool              `json:"featured,omitempty"`
	Parameters         map[int]Parameter `json:"parameters,omitempty"`
	Ranking            int               `json:"ranking,omitempty"`
	SlashID            string            `json:"_id,omitempty"`
	TrainingService    string            `json:"trainingService,omitempty"`
	PredictionService  string            `json:"predictionService,omitempty"`
	ReportService      string            `json:"reportService,omitempty"`
}

// JaqpotEntities structure
type JaqpotEntities struct {
	Total          int
	JaqpotEntities []JaqpotEntity
}

// JaqpotEntity structure
type JaqpotEntity struct {
	ID                 string   `json:"id,omitempty"`
	Meta               MetaInfo `json:"meta,omitempty"`
	OntologicalClasses []string `json:"ontologicalClasses,omitempty"`
	Visible            bool     `json:"visible,omitempty"`
	Temporary          bool     `json:"temporary,omitempty"`
	Featured           bool     `json:"featured,omitempty"`
}

// EntryID structure
type EntryID struct {
	Name      string `json:"name,omitempty"`
	OwnerUUID string `json:"ownerUUID,omitempty"`
	URI       string `json:"URI"`
	Type      string `json:"type,omitempty"`
}

// DataEntry structure
type DataEntry struct {
	EntryID EntryID                `json:"entryId,omitempty"`
	Values  map[string]interface{} `json:"values,omitempty"`
}

// Dataset structure
type Dataset struct {
	Meta               MetaInfo      `json:"meta,omitempty"`
	OntologicalClasses []string      `json:"ontologicalClasses,omitempty"`
	Visible            bool          `json:"visible,omitempty"`
	Temporary          bool          `json:"temporary,omitempty"`
	Featured           bool          `json:"featured,omitempty"`
	DatasetURI         string        `json:"datasetURI,omitempty"`
	ByModel            string        `json:"byModel,omitempty"`
	DataEntry          []DataEntry   `json:"dataEntry,omitempty"`
	Features           []FeatureInfo `json:"features,omitempty"`
	TotalRows          int           `json:"totalRows,omitempty"`
	TotalColumns       int           `json:"totalColumns,omitempty"`
	SlashID            string        `json:"_id,omitempty"`
	OnTrash            bool          `json:"onTrash,omitempty"`
}

// MetaInfo structure
type MetaInfo struct {
	Identifiers  []string `json:"identifiers,omitempty"`
	Comments     []string `json:"comments,omitempty"`
	Descriptions []string `json:"descriptions,omitempty"`
	Titles       []string `json:"titles,omitempty"`
	Subjects     []string `json:"subjects,omitempty"`
	Publishers   []string `json:"publishers,omitempty"`
	Creators     []string `json:"creators,omitempty"`
	Contributors []string `json:"contributors,omitempty"`
	Audiences    []string `json:"audiences,omitempty"`
	Rights       []string `json:"rights,omitempty"`
	SameAs       []string `json:"sameAs,omitempty"`
	SeeAlso      []string `json:"seeAlso,omitempty"`
	HasSources   []string `json:"hasSources,omitempty"`
	Doi          []string `json:"doi,omitempty"`
	Date         string   `json:"date,omitempty"`
	Picture      string   `json:"picture,omitempty"`
	Markdown     string   `json:"markdown,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Read         []string `json:"read,omitempty"`
	Write        []string `json:"write,omitempty"`
	Execute      []string `json:"execute,omitempty"`
}

// Date structure
type Date struct {
	Year  int
	Month time.Month
	Day   int
	Hour  int
	Min   int
	Sec   int
	Nsec  int
	loc   time.Location
}

// FeatureInfo structure
type FeatureInfo struct {
	Key        string                 `json:"key,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Units      string                 `json:"units,omitempty"`
	Conditions map[string]interface{} `json:"conditions,omitempty"`
	Category   string                 `json:"category,omitempty"`
	URI        string                 `json:"uri,omitempty"`
}

// ErrorReport structure
type ErrorReport struct {
	Meta               MetaInfo     `json:"meta,omitempty"`
	OntologicalClasses []string     `json:"ontologicalClasses,omitempty"`
	Visible            bool         `json:"visible,omitempty"`
	Temporary          bool         `json:"temporary,omitempty"`
	Featured           bool         `json:"featured,omitempty"`
	Code               string       `json:"code,omitempty"`
	Actor              string       `json:"actor,omitempty"`
	Message            string       `json:"message,omitempty"`
	Details            string       `json:"details,omitempty"`
	HTTPStatus         int          `json:"httpStatus,omitempty"`
	Trace              *ErrorReport `json:"trace,omitempty"`
	ID                 string       `json:"id,omitempty"`
}

// Feature structure
type Feature struct {
	Meta               MetaInfo `json:"meta,omitempty"`
	OntologicalClasses []string `json:"ontologicalClasses,omitempty"`
	Visible            bool     `json:"visible,omitempty"`
	Temporary          bool     `json:"temporary,omitempty"`
	Featured           bool     `json:"featured,omitempty"`
	Units              string   `json:"units,omitempty"`
	PredictorFor       string   `json:"predictorFor,omitempty"`
	AdmissibleValues   []string `json:"admissibleValues,omitempty"`
	ID                 string   `json:"id,omitempty"`
	SlashID            string   `json:"_id,omitempty"`
}

// Model structure
type Model struct {
	Meta                 MetaInfo               `json:"meta,omitempty"`
	OntologicalClasses   []string               `json:"ontologicalClasses,omitempty"`
	Visible              bool                   `json:"visible,omitempty"`
	Temporary            bool                   `json:"temporary,omitempty"`
	Featured             bool                   `json:"featured,omitempty"`
	DependentFeatures    []string               `json:"dependentFeatures,omitempty"`
	IndependentFeatures  []string               `json:"independentFeatures,omitempty"`
	PredictedFeatures    []string               `json:"predictedFeatures,omitempty"`
	Reliability          float32                `json:"reliability,omitempty"`
	DatasetURI           string                 `json:"datasetUri,omitempty"`
	Parameters           map[string]interface{} `json:"parameters,omitempty"`
	Algorithm            Algorithm              `json:"algorithm,omitempty"`
	Bibtex               BibTeX                 `json:"bibtex,omitempty"`
	ActualModel          interface{}            `json:"actualModel,omitempty"`
	PmmlModel            interface{}            `json:"pmmlModel,omitempty"`
	AdditionalInfo       interface{}            `json:"additionalInfo,omitempty"`
	PmmlTransformations  string                 `json:"pmmlTransformations,omitempty"`
	DoaModel             string                 `json:"doaModel,omitempty"`
	TransformationModels []string               `json:"transformationModels,omitempty"`
	LinkedModels         []string               `json:"linkedModels,omitempty"`
	ID                   string                 `json:"id,omitempty"`
	SlashID              string                 `json:"_id,omitempty"`
	OnTrash              bool                   `json:"onTrash,omitempty"`
}

// BibTeX structure
type BibTeX struct {
	Meta               MetaInfo `json:"meta,omitempty"`
	OntologicalClasses []string `json:"ontologicalClasses,omitempty"`
	Visible            bool     `json:"visible,omitempty"`
	Temporary          bool     `json:"temporary,omitempty"`
	Featured           bool     `json:"featured,omitempty"`
	Author             string   `json:"author,omitempty"`
	Title              string   `json:"title,omitempty"`
	BookTitle          string   `json:"bookTitle,omitempty"`
	School             string   `json:"school,omitempty"`
	Chapter            string   `json:"chapter,omitempty"`
	Copyright          string   `json:"copyright,omitempty"`
	Edition            string   `json:"edition,omitempty"`
	Editor             string   `json:"editor,omitempty"`
	Crossref           string   `json:"crossref,omitempty"`
	Address            string   `json:"address,omitempty"`
	Year               string   `json:"year,omitempty"`
	Pages              string   `json:"pages,omitempty"`
	Volume             string   `json:"volume,omitempty"`
	Number             string   `json:"number,omitempty"`
	Journal            string   `json:"journal,omitempty"`
	Isbn               string   `json:"isbn,omitempty"`
	Issn               string   `json:"issn,omitempty"`
	Keywords           string   `json:"keywords,omitempty"`
	Key                string   `json:"key,omitempty"`
	Annotation         string   `json:"annotation,omitempty"`
	Series             string   `json:"series,omitempty"`
	URL                string   `json:"url,omitempty"`
	BibType            string   `json:"bibType,omitempty"`
	Publisher          string   `json:"publisher,omitempty"`
	ID                 string   `json:"id,omitempty"`
	Abstract           string   `json:"abstract,omitempty"`
}

// Models structure
type Models struct {
	Total  int
	Models []Model
}

// Parameter structure
type Parameter struct {
	Name          string        `json:"name,omitempty"`
	Value         interface{}   `json:"value,omitempty"`
	Scope         string        `json:"scope,omitempty"`
	AllowedValues []interface{} `json:"allowedValues,omitempty"`
	MinValue      interface{}   `json:"minValue,omitempty"`
	MaxValue      interface{}   `json:"maxValue,omitempty"`
	MinArraySize  int           `json:"minArraySize,omitempty"`
	MaxArraySize  int           `json:"maxArraySize,omitempty"`
	Description   string        `json:"description,omitempty"`
}

// Prediction structure
type Prediction struct {
	ModelID     string                   `json:"modelId,omitempty"`
	DatasetID   string                   `json:"datasetId,omitempty"`
	Data        []map[string]interface{} `json:"data,omitempty"`
	Predictions []map[string]interface{} `json:"predictions,omitempty"`
}

// Substance structure
type Substance struct {
	URI       string `json:"URI,omitempty"`
	Name      string `json:"name,omitempty"`
	OwnerUUID string `json:"ownerUUID,omitempty"`
}

// Task structure
type Task struct {
	ID                  string      `json:"id,omitempty"`
	Meta                MetaInfo    `json:"meta,omitempty"`
	OntologicalClasses  []string    `json:"ontologicalClasses,omitempty"`
	Visible             bool        `json:"visible,omitempty"`
	Temporary           bool        `json:"temporary,omitempty"`
	Featured            bool        `json:"featured,omitempty"`
	SlashID             string      `json:"_id,omitempty"`
	ResultURI           string      `json:"resultUri,omitempty"`
	Result              string      `json:"result,omitempty"`
	HasStatus           string      `json:"hasStatus,omitempty"`
	PercentageCompleted float32     `json:"percentageCompleted,omitempty"`
	ErrorReport         ErrorReport `json:"errorReport,omitempty"`
	HTTPStatus          int         `json:"httpStatus,omitempty"`
	Duration            float64     `json:"duration,omitempty"`
	Type                string      `json:"type,omitempty"`
}

// Trained structure
type Trained struct {
	RawModel            interface{} `json:"rawModel,omitempty"`
	PmmlModel           interface{} `json:"pmmlModel,omitempty"`
	AdditionalInfo      interface{} `json:"additionalInfo,omitempty"`
	DependentFeatures   []string    `json:"dependentFeatures,omitempty"`
	IndependentFeatures []string    `json:"independentFeatures,omitempty"`
	PredictedFeatures   []string    `json:"predictedFeatures,omitempty"`
	Runtime             []string    `json:"runtime,omitempty"`
	ImplementedWith     []string    `json:"implementedWith,omitempty"`
	Title               []string    `json:"title,omitempty"`
	Description         []string    `json:"description,omitempty"`
	Algorithm           []string    `json:"algorithm,omitempty"`
	Batched             bool        `json:"batched,omitempty"`
}

// Doa structure
type Doa struct {
	Meta      MetaInfo    `json:"meta,omitempty"`
	ModelID   string      `json:"modelId,omitempty"`
	DoaMatrix [][]float32 `json:"doaMatrix,omitempty"`
	AValue    float32     `json:"aValue,omitempty"`
}
