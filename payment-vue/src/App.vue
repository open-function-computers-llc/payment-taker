<template>
	<div id="app">
		<div class="container">
			<a class="home-logo" href="/">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="312.74"
					height="168.48"
					viewBox="0 0 312.74 168.48"
				>
					<g
						id="Group_18"
						data-name="Group 18"
						transform="translate(-648.26 -347.44)"
					>
						<text
							id="open_"
							transform="translate(709 425)"
							fill="white"
							font-size="76"
							font-family="Montserrat-ExtraBold, Montserrat"
							font-weight="800"
						>
							<tspan x="0" y="0">open_</tspan>
						</text>
						<text
							id="function"
							transform="translate(701 489)"
							fill="white"
							font-size="57"
							font-family="Montserrat-ExtraBold, Montserrat"
							font-weight="800"
							letter-spacing="0.002em"
						>
							<tspan x="0" y="0">function</tspan>
						</text>
						<path
							id="Path_12"
							data-name="Path 12"
							d="M41.94-65.88q0,7.38-2.34,11.16a12.965,12.965,0,0,1-7.56,5.4,12.965,12.965,0,0,1,7.56,5.4q2.34,3.78,2.34,11.16V6.66q0,14.04,13.32,14.04h4.5V34.92H51.84q-13.5,0-20.34-6.75T24.66,8.28v-41.4q0-9-8.1-9h-6.3v-14.4h6.3q8.1,0,8.1-9v-41.4q0-13.14,6.84-19.89t20.34-6.75h7.92v14.22h-4.5q-13.32,0-13.32,14.04Z"
							transform="translate(638 481)"
							fill="white"
						/>
					</g>
				</svg>
			</a>

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
				v-bind:email="this.companyEmail"
			/>
		</div>
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
			companyEmail: "",
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
			this.companyEmail = payload.email;
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
	height: 100vh;
	width: 100%;
	background: #1f1f1f;
}
.home-logo svg {
	width: 120px;
	height: 100px;
}
.logo-home:hover {
	text-decoration: none;
	color: #222222;
}
.container {
	padding-top: 1em;
}
.app-card {
	max-width: 600px;
	width: 100%;
	margin: 100px auto 0;
	background-size: cover;
	background-repeat: no-repeat;
	background-color: #6c63ff;
	border-radius: 5px;
	box-shadow: 3px 3px 30px 0px rgba(0, 0, 0, 0.3);
}
#card-element {
	background: white;
}
.blob-body {
	max-width: 600px;
	width: 100%;
	margin: 0 auto;
	padding: 4.5rem 1.25rem;
	color: white;
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
	color: white;
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
