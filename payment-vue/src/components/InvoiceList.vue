<template>
	<div class="app-container">
		<div class="app-card">
			<div class="blob-body">
				<h1 class="mb-3 text-center">Pay Invoice</h1>

				<div class="invoice-form">
					<div v-show="hideCompanyEmail === false">
						<label>Company</label>
						<div class="input-group mb-3">
							<input
								id="company"
								class="form-control"
								type="text"
								v-model="company"
								ref="company"
								autofocus
							/>
						</div>
						<label>Email</label>
						<div class="input-group mb-3">
							<input
								id="email"
								name="email"
								class="form-control"
								type="email"
								v-model="email"
								ref="email"
							/>
						</div>
					</div>
					<label
						>Invoice Number
						<span class="tooltip">
							<sup><i class="fa fa-info-circle"></i></sup>
							<span class="tooltiptext">
								If you are paying for multiple invoices, this
								can be left blank
							</span>
						</span></label
					>
					<div class="input-group mb-3">
						<input
							id="invoice"
							name="invoice"
							class="form-control"
							type="text"
							v-model="newInvoiceNumber"
							ref="invoiceNumber"
						/>
					</div>
					<label for="amount">Amount Due</label>
					<div class="input-group mb-3">
						<div class="input-group-prepend">
							<span class="input-group-text">$</span>
						</div>
						<input
							class="form-control"
							id="amount"
							type="text"
							v-model="newInvoiceAmount"
							ref="invoiceAmount"
						/>
					</div>
					<div class="text-center">
						<button @click="addInvoice" class="btn mb-5">
							Add Invoice
						</button>
					</div>
					<table v-if="invoices.length" class="mb-3">
						<tr
							v-for="(invoice, index) in invoices"
							:key="'invoice' + index"
						>
							<td>Invoice # {{ invoice.number }}</td>
							<td class="text-right">
								${{ parseFloat(invoice.amount).toFixed(2) }}
							</td>
						</tr>
						<tr style="border-bottom: 1px solid #ccc">
							<td>
								<span class="fee-line"
									>Online Transaction Fee</span
								>
								<span class="tooltip">
									<sup><i class="fa fa-info-circle"></i></sup>
									<span class="tooltiptext">
										Please send a check by mail to avoid the
										online transaction fee.
									</span>
								</span>
							</td>
							<td class="text-right">${{ total.fee }}</td>
						</tr>
						<tr>
							<td>Total</td>
							<td class="text-right">${{ total.totalCharge }}</td>
						</tr>
					</table>

					<div class="text-center">
						<button
							class="btn w-100"
							v-if="invoices.length"
							@click="$emit('ready-to-pay', paymentInfo)"
						>
							Pay
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	data() {
		return {
			newInvoiceNumber: "",
			newInvoiceAmount: 0,
			show: false,
			company: "",
			email: "",
			fee: "",
			hideCompanyEmail: false,
			invoices: [],
		};
	},
	computed: {
		// Calculate stripe fee from amount
		// so we can charge stripe fee to customers
		total() {
			var amt = 0;
			var total = 0;
			var stripeFee = { Percent: 2.9, Fixed: 0.3 };
			var chargedFee = 0;

			this.invoices.map((i) => {
				amt += parseFloat(i.amount);
			});

			total =
				(amt + parseFloat(stripeFee.Fixed)) /
				(1 - parseFloat(stripeFee.Percent) / 100);

			chargedFee = total - amt;

			return {
				amt: amt.toFixed(2),
				fee: chargedFee.toFixed(2),
				totalCharge: total.toFixed(2),
			};
		},

		// Add amount, invoice info and company name to the payload
		// this is posted to the PaymentIntent API in App.vue
		paymentInfo() {
			return {
				amount: this.total,
				invoices: this.invoices,
				email: this.email,
				company: this.company,
			};
		},
	},
	methods: {
		addInvoice() {
			this.invoices.push({
				number: this.newInvoiceNumber,
				amount: this.newInvoiceAmount.toString(),
			});
			this.hideCompanyEmail = true;
			this.newInvoiceNumber = "";
			this.newInvoiceAmount = 0;
			this.$refs.invoiceNumber.focus();
		},
	},
};
</script>

<style>
.invoice-form {
	max-width: 350px;
	margin: 0 auto;
}

.tooltip {
	position: relative;
	display: inline-block;
	opacity: 1;
}

.tooltip .tooltiptext {
	visibility: hidden;
	width: 220px;
	background-color: #222222;
	color: #fff;
	border-radius: 6px;
	padding: 5px 10px;

	/* Position the tooltip */
	position: absolute;
	z-index: 1;
}

.tooltip:hover .tooltiptext {
	visibility: visible;
}
</style>
