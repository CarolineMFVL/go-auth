# JWT Documentation

## Tokens

### Definitions

The GenerateJWT function creates a signed JWT token containing the user ID, email, expiration, and token type (access or refresh):

Access token: Short-lived (15 min), used for authenticating API requests.
Refresh token: Long-lived (24h), used to obtain new access tokens.

### JWT Validation

The ValidateJWT function checks the validity and integrity of a JWT token (typically received in the Authorization header):

Verifies the signature using the secret key.
Decodes the claims (payload) from the token.
Returns an error if the token is expired, malformed, or tampered with.

### Usage in the Login Handler

When a user logs in successfully, two tokens are generated and returned:

### Security Best Practices

Never expose the secret key (jwtKey) in your source code; always use an environment variable.
Always check the token type (token_type) when using tokens (e.g., only accept access tokens for protected resources).
Renew access tokens using the refresh token.
Use HTTPS to protect tokens in transit.

### Example Client Usage

Send the access token in the header:

`Authorization: Bearer <access_token>`
