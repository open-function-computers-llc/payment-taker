# Payment taker program

This app is here to allow for people to accept online payments via Stripe.

## Oddities

### Env mess

Vue has it's own `env` and so does go. That's not that fun. Make sure that you put your stripe testing key in both `.env` and `./payment-vue/.env`. When the vue app gets built for prod, make sure you reference the stripe key from the Go env file. To do this, make sure you have a file named `.env.production` in the vue app folder, and that the key `VUE_APP_STRIPE_PUB_KEY` is set to this: `{{ .StripePublicKey }}`

### Node dev server

If you want to run the node dev server, from the sub-folder /payment-vue run `npm run serve` and make sure you visit the app at the sub folder address from the node server. On my machine, that's this: `http://localhost:8080/static/`

### To Build Go

From the root run `./build.sh`
