package lib


// AuthError openid complinent & oauth2.0 error codes
type AuthError string

const (
	/* InteractionRequired:
	The Authorization Server requires End-User interaction of some form to proceed. 
	This error MAY be returned when the prompt parameter value in the Authentication Request is none, 
	but the Authentication Request cannot be completed without displaying a user interface for End-User interaction.
	*/
	InteractionRequired   AuthError = "interaction_required"
	
	/* LoginRequired:
	The Authorization Server requires End-User authentication. 
	This error MAY be returned when the prompt parameter value in the Authentication Request is none, 
	but the Authentication Request cannot be completed without displaying a user interface for End-User authentication.
	*/
	LoginRequired AuthError = "login_required"
	
	/* AccountSelectionRequired:
	The End-User is REQUIRED to select a session at the Authorization Server.
	The End-User MAY be authenticated at the Authorization Server with different associated accounts,
	but the End-User did not select a session. This error MAY be returned when the prompt parameter value in the Authentication Request is none,
	but the Authentication Request cannot be completed without displaying a user interface to prompt for a session to use.
	*/
	AccountSelectionRequired AuthError = "account_selection_required"

	/* ConsentRequired:
	The Authorization Server requires End-User consent.
	This error MAY be returned when the prompt parameter value in the Authentication Request is none,
	but the Authentication Request cannot be completed without displaying a user interface for End-User consent.
	*/
	ConsentRequired AuthError = "consent_required"

	/* InvalidRequestURI:
	The request_uri in the Authorization Request returns an error or contains invalid data.
	*/
	InvalidRequestURI AuthError = "invalid_request_uri"

	/*InvalidRequestObject:
	The request parameter contains an invalid Request Object
	*/
	InvalidRequestObject AuthError = "invalid_request_object"

	/* RequestNotSupported:
	The OP does not support use of the request parameter
	*/
	RequestNotSupported AuthError = "request_not_supported"

	/* RequestUriNotSupported:
	The OP does not support use of the request_uri parameter 
	*/
	RequestURINotSupported AuthError = "request_uri_not_supported"

	/* RegistrationNotSupported:
	The OP does not support use of the registration parameter
	*/
	RegistrationNotSupported AuthError = "registration_not_supported"

	/* Invalid Request
	Oauth2.0 This error occurs when there is a missing parameter that includes multiple credentials,
	unsupported parameter value.
	*/
	InvalidRequest AuthError = "invalid_request"

	/* Unauthorized Client
	Oauth2.0 The unauthorized client is not allowed to access the authorization grant type.
	*/
	UnauthorizedClient AuthError = "unauthorized_client"

	/* Invalid Client
	Bad client auth
	*/
	InvalidClient AuthError = "invalid_client"

	/* Invalid Token
	*/
	InvalidToken AuthError = "invalid_token"
)
