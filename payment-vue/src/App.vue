<template>
  <div id="app">
    <InvoiceList v-if="state === 'gatherInvoices'" v-on:ready-to-pay="collectPayment" />
    <PayForm v-if="state === 'readyToPay'" v-bind:intent="paymentIntentID" />
  </div>
</template>

<script>
import axios from "axios"
import InvoiceList from './components/InvoiceList.vue'
import PayForm from './components/PayForm.vue'

export default {
  name: 'App',
  data() {
    return {
      state: "gatherInvoices",
      amount: 0,
      paymentIntentID: null,
    }
  },
  components: {
    InvoiceList,
    PayForm
  },
  methods: {
    collectPayment(payload) {
      this.state = "readyToPay"
      this.amount = payload.amount

      axios({
        method: 'post',
        url: process.env.VUE_APP_API_DOMAIN + '/intent',
        data: {
          invoices: payload.invoices,
          amount: Math.round(this.amount * 100) // pennies
        }
      }).then((response) => {
        console.log(response)
        this.paymentIntentID = response.data.intentID
      })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
