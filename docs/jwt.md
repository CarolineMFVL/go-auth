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

### Refresh Token Storage and Blacklisting

Store refresh tokens securely if you want to support logout or token blacklisting:

- **Server-side storage:** Save refresh tokens in a secure database table associated with the user. This allows you to revoke (blacklist) tokens on logout or in case of compromise.
- **Token rotation:** Issue a new refresh token each time one is used, and invalidate the previous one.
- **Blacklist mechanism:** Maintain a blacklist of revoked tokens and check it on each refresh request.
- **Expiration:** Always set an expiration time for refresh tokens and periodically clean up expired tokens from storage.

**Best practice:** Never store refresh tokens in client-side insecure storage (like localStorage). Prefer HTTP-only cookies or server-side storage for maximum security.
