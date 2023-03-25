# Punkz documentation

## API

### GET /api/auth?provider=discord

Redirects to the Discord OAuth2 login page.

### GET /api/auth/callback

Return from the Discord OAuth2 login page.
Automatically redirects to the frontend.

### GET /api/user

Returns the user's data.

Return body: 

```json
{
	"avatar": "",
	"email": "exiletlg@gmail.com",
	"id": "779887965107126303",
	"name": "fiksik"
}
```

### POST /api/wallet

Creates a new wallet.

Headers: 
```json
    content-type: application/json
```

Body: 
```json
    {
        "wallet": "solana wallet"
    }
```

### GET /api/ping

Returns a pong.

