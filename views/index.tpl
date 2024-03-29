<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <title>Online Payment Portal | {{ .Client }}</title>

    <link rel="icon" href="/static/favicon.svg">
    <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap" rel="stylesheet">
    <link rel="stylesheet" media="print" href="/static/print.css" />
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
