<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <title>Online Payment Portal | {{ .Client }}</title>

    <link rel="icon" href="/static/favicon.ico">
    <script src="//js.stripe.com/v3/"></script>
    <script>
      var stripePublicKey = "{{ .StripePublicKey }}";
      var stripe = Stripe(stripePublicKey);
    </script>
  <link href="/static/js/app.js" rel="preload" as="script"><link href="/static/js/chunk-vendors.js" rel="preload" as="script"></head>
  <body>
    <noscript>
      <strong>We're sorry but this doesn't work properly without JavaScript enabled. Please enable it to continue.</strong>
    </noscript>
    <div id="app"></div>
    <!-- built files will be auto injected -->
  <script type="text/javascript" src="/static/js/chunk-vendors.js"></script><script type="text/javascript" src="/static/js/app.js"></script></body>
</html>
