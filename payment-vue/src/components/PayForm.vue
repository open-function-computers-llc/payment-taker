<template>
	<div class="app-container">
		<div v-if="state === 'payForm'" class="app-card">
			<div class="blob-body">
				<div class="pay-form">
					<h3 class="details-header">${{ total }}</h3>
					<div :v-if="card" id="card-element" class="mb-2">
						<!-- Create stripe card input elements here -->
					</div>

					<!-- Throw error messages in this element -->
					<div id="card-errors" class="mb-3" role="alert"></div>
					<button id="submit" class="btn w-100" @click="pay">
						Pay
					</button>
				</div>
			</div>
		</div>

		<div v-if="state === 'error'" class="app-card">
			<div class="blob-body">
				<h1 class="text-center">Yikes!</h1>
				<h3 class="text-center">
					There has been an issue processing your request, please try
					again.
				</h3>
			</div>
		</div>

		<div v-if="state === 'thanks'" class="app-card">
			<div class="blob-body">
				<h3 class="text-center mb-4">OPEN FUNCTION</h3>
				<p class="text-center">{{ company }} paid on {{ timestamp }}</p>
				<div
					v-for="(invoice, index) in invoices"
					:key="'invoice' + index"
				>
					<div class="flex">
						<span>Invoice # {{ invoice.number }}</span>
						<span
							>${{ parseFloat(invoice.amount).toFixed(2) }}</span
						>
					</div>
				</div>
				<hr />
				<div class="flex">
					<span>
						Online Transaction Fee
						<span class="tooltip">
							<sup><i class="fa fa-info-circle"></i></sup>
							<span class="tooltiptext">
								Please send a check by mail to avoid the online
								transaction fee.
							</span>
						</span>
					</span>
					<span>${{ fee }}</span>
				</div>
				<div class="flex mb-4">
					<span>Total</span>
					<span>${{ total }}</span>
				</div>
				<p class="text-center">Thank you for your payment!</p>
			</div>
		</div>
	</div>
</template>

<script>
import axios from "axios";

export default {
	data() {
		return {
			complete: false,
			publicKey: window.stripePublicKey,
			card: null,
			state: "payForm",
			spinner: true,
			timestamp: "",
		};
	},
	props: ["secret", "company", "email", "total", "invoices", "fee"],
	mounted() {
		// Get Stripe.js Elements to use in checkout form
		var elements = window.stripe.elements();
		var style = {
			base: {
				fontFamily: '"Montserrat", Helvetica, sans-serif',
				fontSmoothing: "antialiased",
				fontSize: "16px",
			},
		};

		var cardJS = elements.create("card", { style: style });
		cardJS.mount("#card-element");
		cardJS.on("change", function (event) {
			var displayError = document.getElementById("card-errors");
			if (event.error) {
				displayError.textContent = event.error.message;
			} else {
				displayError.textContent = "";
			}
		});

		this.card = cardJS;
	},
	created() {
		setInterval(this.getNow, 1000);
	},
	methods: {
		pay() {
			// client secret is a prop passed from App.vue
			// stripe payment is processed with a token that can either be
			//  created on the front-end or returned by posting to stripe
			// we posted to the intent API which returned the
			// client secret we need to process transaction
			this.spinner = "true";
			window.stripe
				.confirmCardPayment(this.secret, {
					// email: this.email,
					payment_method: {
						card: this.card,
					},
				})
				.then((result) => {
					if (result.error) {
						console.log("Error!", result.error);
						this.state = "error";
					} else {
						this.spinner = false;
						console.log("Success!", result, this.email);
						this.state = "thanks";

						// do ajax request to send reciepts to us and customer
						axios.post(
							process.env.VUE_APP_API_DOMAIN + "/receipt",
							{
								// "result" : result,
								email: this.email,
								total: this.total,
								fee: this.fee,
							}
						);
					}
				});
		},
		getNow: function () {
			const today = new Date();
			const date =
				today.getMonth() +
				1 +
				"-" +
				today.getDate() +
				"-" +
				today.getFullYear();
			this.timestamp = date;
		},
	},
};
</script>

<style>
.pay-form {
	max-width: 420px;
	margin: 0 auto;
}
#card-element {
	padding: 1em 10px;
	border-radius: 5px;
	box-shadow: 0px 0px 18px -5px rgba(0, 0, 0, 0.3);
}
#card-errors {
	font-size: 14px;
	color: rgb(230, 0, 0);
}
.details-header {
	color: #7a3cfc;
	text-align: center;
	font-weight: bold;
	font-size: 26px;
	padding: 0 0 20px;
}
#__lpform_cardnumber_icon {
	display: none;
}
.flex {
	display: flex;
	justify-content: space-between;
}
.bouncingLoader > div,
.bouncingLoader:before,
.bouncingLoader:after {
	display: inline-block;
	width: 13px;
	height: 13px;
	background: hsl(243, 80%, 62%);
	margin-bottom: -5px;
	border-radius: 50%;
	animation: bouncing-loader 0.6s infinite alternate;
}
.bouncingLoader > div,
.bouncingLoader:before,
.bouncingLoader:after {
	content: "";
}
.bouncingLoader > div {
	margin: 0 5px;
}
.bouncingLoader > div {
	animation-delay: 0.2s;
}
.bouncingLoader:after {
	animation-delay: 0.4s;
}

@keyframes bouncing-loader {
	to {
		opacity: 0.1;
		transform: translate3d(0, -16px, 0);
	}
}
</style>
