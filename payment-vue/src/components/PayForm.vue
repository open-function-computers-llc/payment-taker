<template>
    <div>
        <h1>Please give us your payment details:</h1>
        <card class='stripe-card'
            :class='{ complete }'
            :stripe='publicKey'
            :options='stripeOptions'
            @change='complete = $event.complete'
        />
        <button class='pay-with-stripe' @click='pay' :disabled='!complete'>Pay with credit card</button>
    </div>
</template>

<script>
import { Card, createToken } from 'vue-stripe-elements-plus'

export default {
    data() {
        return {
            complete: false,
            publicKey: window.stripePublicKey,
        }
    },
    props: ["intent"],
    components: { Card },
    methods: {
        pay () {
            console.log(this)
            console.log(this.intent)
            // createToken returns a Promise which resolves in a result object with
            // either a token or an error key.
            // See https://stripe.com/docs/api#tokens for the token object.
            // See https://stripe.com/docs/api#errors for the error object.
            // More general https://stripe.com/docs/stripe.js#stripe-create-token.
            createToken().then(data => {
                window.stripe.confirmCardPayment(this.intent, {
                    payment_method: {
                        card: data.token,
                        billing_details: {
                            name: 'Jenny Rosen'
                        }
                    }
                }).then((result) => {
                    if (result.error) {
                        console.log("Error!", result)
                    } else {
                        console.log("Success!", result)
                    }
                })
                console.log(data.token)
            })
        }
    }
}
</script>
