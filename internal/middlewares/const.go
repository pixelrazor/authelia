package middlewares

import (
	"errors"

	"github.com/valyala/fasthttp"
)

var (
	headerAccept        = []byte(fasthttp.HeaderAccept)
	headerContentLength = []byte(fasthttp.HeaderContentLength)

	headerXForwardedProto = []byte(fasthttp.HeaderXForwardedProto)
	headerXForwardedHost  = []byte(fasthttp.HeaderXForwardedHost)
	headerXForwardedFor   = []byte(fasthttp.HeaderXForwardedFor)
	headerXRequestedWith  = []byte(fasthttp.HeaderXRequestedWith)

	headerXForwardedURI    = []byte("X-Forwarded-URI")
	headerXOriginalURL     = []byte("X-Original-URL")
	headerXForwardedMethod = []byte("X-Forwarded-Method")

	headerVary   = []byte(fasthttp.HeaderVary)
	headerAllow  = []byte(fasthttp.HeaderAllow)
	headerOrigin = []byte(fasthttp.HeaderOrigin)

	headerAccessControlAllowCredentials = []byte(fasthttp.HeaderAccessControlAllowCredentials)
	headerAccessControlAllowHeaders     = []byte(fasthttp.HeaderAccessControlAllowHeaders)
	headerAccessControlAllowMethods     = []byte(fasthttp.HeaderAccessControlAllowMethods)
	headerAccessControlAllowOrigin      = []byte(fasthttp.HeaderAccessControlAllowOrigin)
	headerAccessControlMaxAge           = []byte(fasthttp.HeaderAccessControlMaxAge)
	headerAccessControlRequestHeaders   = []byte(fasthttp.HeaderAccessControlRequestHeaders)
	headerAccessControlRequestMethod    = []byte(fasthttp.HeaderAccessControlRequestMethod)

	headerXContentTypeOptions   = []byte(fasthttp.HeaderXContentTypeOptions)
	headerReferrerPolicy        = []byte(fasthttp.HeaderReferrerPolicy)
	headerXFrameOptions         = []byte(fasthttp.HeaderXFrameOptions)
	headerPragma                = []byte(fasthttp.HeaderPragma)
	headerCacheControl          = []byte(fasthttp.HeaderCacheControl)
	headerXXSSProtection        = []byte(fasthttp.HeaderXXSSProtection)
	headerContentSecurityPolicy = []byte(fasthttp.HeaderContentSecurityPolicy)

	headerPermissionsPolicy = []byte("Permissions-Policy")

	headerCrossOriginEmbedderPolicy = []byte("Cross-Origin-Embedder-Policy")
	headerCrossOriginOpenerPolicy   = []byte("Cross-Origin-Opener-Policy")
	headerCrossOriginResourcePolicy = []byte("Cross-Origin-Resource-Policy")
)

var (
	headerValueFalse          = []byte("false")
	headerValueTrue           = []byte("true")
	headerValueMaxAge         = []byte("100")
	headerValueVary           = []byte("Accept-Encoding, Origin")
	headerValueVaryWildcard   = []byte("Accept-Encoding")
	headerValueOriginWildcard = []byte("*")
	headerValueZero           = []byte("0")
	headerValueCSPNone        = []byte("default-src 'none';")

	headerValueNoSniff                 = []byte("nosniff")
	headerValueStrictOriginCrossOrigin = []byte("strict-origin-when-cross-origin")
	headerValueSAMEORIGIN              = []byte("SAMEORIGIN")
	headerValueNoCache                 = []byte("no-cache")
	headerValueNoStore                 = []byte("no-store")
	headerValueXSSModeBlock            = []byte("1; mode=block")
	headerValueCohort                  = []byte("interest-cohort=()")

	headerValueRequireCORP = []byte("require-corp")
	headerValueSameOrigin  = []byte("same-origin")
	headerValueUnsafeNone  = []byte("unsafe-none")
)

var (
	protoHTTPS = []byte("https")
	protoHTTP  = []byte("http")

	// UserValueKeyBaseURL is the User Value key where we store the Base URL.
	UserValueKeyBaseURL = []byte("base_url")

	headerSeparator = []byte(", ")
)

const (
	headerValueXRequestedWithXHR = "XMLHttpRequest"
	contentTypeApplicationJSON   = "application/json"
	contentTypeTextHTML          = "text/html"
)

var okMessageBytes = []byte("{\"status\":\"OK\"}")

const (
	messageOperationFailed                      = "Operation failed"
	messageIdentityVerificationTokenAlreadyUsed = "The identity verification token has already been used"
	messageIdentityVerificationTokenHasExpired  = "The identity verification token has expired"
)

var protoHostSeparator = []byte("://")

var errPasswordPolicyNoMet = errors.New("the supplied password does not met the security policy")
