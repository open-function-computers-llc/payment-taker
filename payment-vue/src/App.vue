<template>
    <div id="app">
        <InvoiceList
            v-if="state === 'gatherInvoices'"
            v-on:ready-to-pay="collectPayment"
        />
        <PayForm
            v-if="state === 'readyToPay'"
            v-bind:intent="paymentIntentID"
        />
    </div>
</template>

<script>
import axios from "axios";
import InvoiceList from "./components/InvoiceList.vue";
import PayForm from "./components/PayForm.vue";
import Vue from "vue";

import BootstrapVue from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.use(BootstrapVue);

export default {
    name: "App",
    data() {
        return {
            state: "gatherInvoices",
            amount: 0,
            paymentIntentID: null,
        };
    },
    components: {
        InvoiceList,
        PayForm,
    },
    methods: {
        collectPayment(payload) {
            this.state = "readyToPay";
            this.amount = payload.amount;
            axios({
                method: "post",
                url: process.env.VUE_APP_API_DOMAIN + "/intent",
                data: {
                    invoices: payload.invoices,
                    amount: Math.round(this.amount * 100), // pennies
                },
            }).then((response) => {
                this.paymentIntentID = response.data.paymentIntentID;
            });
        },
    },
};
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #f9a673;
}
</style>
