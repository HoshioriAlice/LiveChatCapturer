// Generated by https://quicktype.io

package main

type LiveChat struct {
	Timing    Timing   `json:"timing"`    
	Response  Response `json:"response"`  
	Endpoint  Endpoint `json:"endpoint"`  
	URL       string   `json:"url"`       
	Csn       string   `json:"csn"`       
	XsrfToken string   `json:"xsrf_token"`
}

type Endpoint struct {
	CommandMetadata EndpointCommandMetadata `json:"commandMetadata"`
	URLEndpoint     URLEndpoint             `json:"urlEndpoint"`    
}

type EndpointCommandMetadata struct {
	WebCommandMetadata URLEndpoint `json:"webCommandMetadata"`
}

type URLEndpoint struct {
	URL string `json:"url"`
}

type Response struct {
	ResponseContext      ResponseContext      `json:"responseContext"`     
	ContinuationContents ContinuationContents `json:"continuationContents"`
	TrackingParams       string               `json:"trackingParams"`      
}

type ContinuationContents struct {
	LiveChatContinuation LiveChatContinuation `json:"liveChatContinuation"`
}

type LiveChatContinuation struct {
	Continuations  []Continuation `json:"continuations"` 
	Actions        []Action       `json:"actions"`       
	TrackingParams TrackingParams `json:"trackingParams"`
}

type Action struct {
	AddChatItemAction AddChatItemAction `json:"addChatItemAction"`
}

type AddChatItemAction struct {
	Item     Item   `json:"item"`    
	ClientID string `json:"clientId"`
}

type Item struct {
	LiveChatTextMessageRenderer LiveChatTextMessageRenderer `json:"liveChatTextMessageRenderer"`
}

type LiveChatTextMessageRenderer struct {
	Message                  Title               `json:"message"`                 
	AuthorName               Title               `json:"authorName"`              
	AuthorPhoto              AuthorPhoto         `json:"authorPhoto"`             
	ContextMenuEndpoint      ContextMenuEndpoint `json:"contextMenuEndpoint"`     
	ID                       string              `json:"id"`                      
	TimestampUsec            string              `json:"timestampUsec"`           
	AuthorBadges             []AuthorBadge       `json:"authorBadges"`            
	AuthorExternalChannelID  string              `json:"authorExternalChannelId"` 
	ContextMenuAccessibility Accessibility       `json:"contextMenuAccessibility"`
}

type AuthorBadge struct {
	LiveChatAuthorBadgeRenderer LiveChatAuthorBadgeRenderer `json:"liveChatAuthorBadgeRenderer"`
}

type LiveChatAuthorBadgeRenderer struct {
	CustomThumbnail CustomThumbnail `json:"customThumbnail"`
	Tooltip         Label           `json:"tooltip"`        
	Accessibility   Accessibility   `json:"accessibility"`  
}

type Accessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}

type AccessibilityData struct {
	Label Label `json:"label"`
}

type CustomThumbnail struct {
	Thumbnails []URLEndpoint `json:"thumbnails"`
}

type Title struct {
	SimpleText string `json:"simpleText"`
}

type AuthorPhoto struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	URL    string `json:"url"`   
	Width  int64  `json:"width"` 
	Height int64  `json:"height"`
}

type ContextMenuEndpoint struct {
	ClickTrackingParams             TrackingParams                     `json:"clickTrackingParams"`            
	CommandMetadata                 ContextMenuEndpointCommandMetadata `json:"commandMetadata"`                
	LiveChatItemContextMenuEndpoint LiveChatItemContextMenuEndpoint    `json:"liveChatItemContextMenuEndpoint"`
}

type ContextMenuEndpointCommandMetadata struct {
	WebCommandMetadata WebCommandMetadata `json:"webCommandMetadata"`
}

type WebCommandMetadata struct {
	IgnoreNavigation bool `json:"ignoreNavigation"`
}

type LiveChatItemContextMenuEndpoint struct {
	Params string `json:"params"`
}

type Continuation struct {
	TimedContinuationData TimedContinuationData `json:"timedContinuationData"`
}

type TimedContinuationData struct {
	TimeoutMS           int64          `json:"timeoutMs"`          
	Continuation        string         `json:"continuation"`       
	ClickTrackingParams TrackingParams `json:"clickTrackingParams"`
}

type ResponseContext struct {
	ServiceTrackingParams           []ServiceTrackingParam          `json:"serviceTrackingParams"`          
	WebResponseContextExtensionData WebResponseContextExtensionData `json:"webResponseContextExtensionData"`
}

type ServiceTrackingParam struct {
	Service string  `json:"service"`
	Params  []Param `json:"params"` 
}

type Param struct {
	Key   string `json:"key"`  
	Value string `json:"value"`
}

type WebResponseContextExtensionData struct {
	YtConfigData   YtConfigData   `json:"ytConfigData"`  
	FeedbackDialog FeedbackDialog `json:"feedbackDialog"`
}

type FeedbackDialog struct {
	PolymerOptOutFeedbackDialogRenderer PolymerOptOutFeedbackDialogRenderer `json:"polymerOptOutFeedbackDialogRenderer"`
}

type PolymerOptOutFeedbackDialogRenderer struct {
	Title         Title      `json:"title"`        
	Subtitle      Disclaimer `json:"subtitle"`     
	Options       []Option   `json:"options"`      
	Disclaimer    Disclaimer `json:"disclaimer"`   
	DismissButton Button     `json:"dismissButton"`
	SubmitButton  Button     `json:"submitButton"` 
	CloseButton   Button     `json:"closeButton"`  
	CancelButton  Button     `json:"cancelButton"` 
}

type Button struct {
	ButtonRenderer ButtonRenderer `json:"buttonRenderer"`
}

type ButtonRenderer struct {
	Style      string `json:"style"`         
	Size       string `json:"size"`          
	IsDisabled bool   `json:"isDisabled"`    
	Text       *Title `json:"text,omitempty"`
	Icon       *Icon  `json:"icon,omitempty"`
}

type Icon struct {
	IconType string `json:"iconType"`
}

type Disclaimer struct {
	Runs []Run `json:"runs"`
}

type Run struct {
	Text               string    `json:"text"`                        
	NavigationEndpoint *Endpoint `json:"navigationEndpoint,omitempty"`
}

type Option struct {
	PolymerOptOutFeedbackOptionRenderer     *PolymerOptOutFeedbackOptionRenderer     `json:"polymerOptOutFeedbackOptionRenderer,omitempty"`    
	PolymerOptOutFeedbackNullOptionRenderer *PolymerOptOutFeedbackNullOptionRenderer `json:"polymerOptOutFeedbackNullOptionRenderer,omitempty"`
}

type PolymerOptOutFeedbackNullOptionRenderer struct {
	Description Title `json:"description"`
}

type PolymerOptOutFeedbackOptionRenderer struct {
	OptionKey           string `json:"optionKey"`          
	Description         Title  `json:"description"`        
	ResponsePlaceholder Title  `json:"responsePlaceholder"`
}

type YtConfigData struct {
	Csn          string `json:"csn"`         
	VisitorData  string `json:"visitorData"` 
	SessionIndex int64  `json:"sessionIndex"`
}

type Timing struct {
	Info Info `json:"info"`
}

type Info struct {
	St int64 `json:"st"`
}

type Label string
const (
	新会员 Label = "新会员"
	评论操作 Label = "评论操作"
)

type TrackingParams string
const (
	CAEQl98BIhMIhYf9R9OS4AIVToTECh0M5Wr TrackingParams = "CAEQl98BIhMIhYf9r9OS4AIVToTECh0M5wr_"
	CAEQl98BIhMItIekrdOS4AIV2MTECh3XXQJc TrackingParams = "CAEQl98BIhMItIekrdOS4AIV2MTECh3xXQJc"
)
