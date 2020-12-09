<template>
    <div id="app">
        <div class="container">
            <a class="logo-home" href="/">
                {of
            </a>
        </div>
        <InvoiceList
            v-if="state === 'gatherInvoices'"
            v-on:ready-to-pay="collectPayment"
        />
        <PayForm
            v-if="state === 'readyToPay'"
            v-bind:secret="this.client_secret"
            v-bind:company="this.company"
            v-bind:total="this.total"
            v-bind:fee="this.fee"
            v-bind:invoices="this.invoices"
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
            fee: 0,
            company: "",
            invoices: [],
            client_scret: "",
        };
    },
    components: {
        InvoiceList,
        PayForm,
    },
    methods: {
        collectPayment(payload) {
            this.total = payload.amount.totalCharge;
            this.fee = payload.amount.fee;
            this.company = payload.company;
            this.email = payload.email;
            this.invoices = payload.invoices;
            // posts amount, invoice info and company name to
            // stripe payment intent API, returns a client_secret
            // that is handled in PayForm to complete transaction
            axios({
                method: "post",
                url: process.env.VUE_APP_API_DOMAIN + "/intent",
                data: {
                    company: payload.company,
                    email: payload.email,
                    invoices: payload.invoices,
                    amount: Math.round(this.total * 100), // pennies
                },
            }).then((result) => {
                // client_secret is used in PayForm
                this.client_secret = result.data.intentID;
                this.state = "readyToPay";
            });
        },
    },
};
</script>

<style>
#app {
    font-family: Montserrat, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    background: #f9a673;
    height: 100vh;
}

.app-container {
    width: 100%;
    display: flex;
    height: calc(100vh - 79px);
    align-items: center;
    justify-content: center;
}
.logo-home {
    font-weight: bold;
    color: #7a3cfc;
    font-size: 42px;
}
.logo-home:hover {
    text-decoration: none;
    color: #222222;
}
.container {
    padding-top: 1em;
}
.app-card {
    max-width: 500px;
    width: 100%;
    margin: 0 auto;
    background-color: white;
    border-radius: 5px;
    box-shadow: 3px 3px 30px 0px rgba(0, 0, 0, 0.3);
}
.card-body {
    max-width: 500px;
    width: 100%;
    margin: 0 auto;
    padding: 3rem 1.25rem;
}

#card-element {
    max-width: 420px;
}

.btn {
    background: #7a3cfc;
    color: white;
    border-radius: 10px;
    font-weight: bold;
}
.btn:hover {
    color: white;
}
label {
    display: block;
    font-weight: 500;
    font-size: 14px;
    margin-bottom: 0;
}
table {
    width: 100%;
}
tr {
    line-height: 32px;
}
.mb-3 {
    margin-bottom: 1.1rem !important;
}
</style>
