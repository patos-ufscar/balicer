package models

type HttpRequestFrame struct {
	Method									string
	RequestURI								string
	HTTPVersion								string
	RequestHeaders							HttpRequestHeaders
	// GeneralHeaders
	// RepresentationHeaders
	RequestBody								[]byte
}

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
type HttpRequestHeaders struct {
	Accept									string
	// AcceptCH								string
    AcceptCharset							string
    AcceptEncoding							string
    AcceptLanguage							string
	// AcceptPatch								string
	// AcceptPost								string
	// AcceptRanges							string

	// AccessControlAllowCredemtials			string
	// AccessControlAllowHeaders				string
	// AccessControlAllowMethods				string
	// AccessControlAllowOrigin				string
	// AccessControlExposeHeaders				string
	// AccessControlMaxAge						string
	// AccessControlRequestHeaders				string
	// AccessControlRequestMethod				string

	// Age										string
	// Allow									string
	// AltSvc									string
	// AltUsed									string

    Authorization							string

	// CacheControl							string
	// ClearSiteData							string
	// Connection								string

	// ContentDigest							string
	// ContentDisposition						string
	// ContentDPR								string
	ContentEncoding							string
	ContentLanguage							string
	ContentLength							uint64
	ContentLocation							string
	// ContentRange							string
	// ContentSecurityPolicy					string
	// ContentSecurityPolicyReportOnly			string
	ContentType								string
	Cookie									string
	// STOPPED HERE!!

	Origin									string
    Host									string
    UserAgent								string
}
