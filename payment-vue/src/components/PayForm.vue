<template>
    <div>
        <h1>Payment Details:</h1>
        <card class='stripe-card'
            :class='{ complete }'
            :stripe='publicKey'
            @change='complete = $event.complete'
        />
        <button class='pay-with-stripe btn w-100' @click='pay' :disabled='!complete'>Submit</button>
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
            })
        }
    }
}
</script>
