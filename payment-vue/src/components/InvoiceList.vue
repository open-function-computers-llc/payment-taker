<template>
    <div class="container">
        <div class="invoice-card">
            <div class="card-body">
                <h1 class="mb-4 text-center">Pay Invoice</h1>
                <label for="company">Company</label>
                <div class="input-group mb-4">
                    <input
                        id="company"
                        class="form-control"
                        type="text"
                        v-model="company"
                        ref="company"
                        autofocus
                        required
                    />
                </div>
                <label for="invoice-number">Invoice Number</label>
                <div class="input-group mb-4">
                    <input
                        id="invoice-number"
                        class="form-control"
                        type="text"
                        v-model="newInvoiceNumber"
                        ref="invoiceNumber"
                        autofocus
                        required
                    />
                </div>
                <label for="amount">Amount Due</label>
                <div class="input-group mb-4">
                    <div class="input-group-prepend">
                        <span class="input-group-text">$</span>
                    </div>
                    <input
                        class="form-control"
                        id="amount"
                        type="number"
                        v-model="newInvoiceAmount"
                        ref="invoiceAmount"
                        required
                    />
                </div>
                <div class="text-center">
                    <button @click="addInvoice" class="btn mb-5">
                        Add Invoice
                    </button>
                </div>
                <table v-if="invoices.length">
                    <tr
                        v-for="(invoice, index) in invoices"
                        :key="'invoice' + index"
                    >
                        <td>Invoice #{{ invoice.number }}</td>
                        <td class="text-right">${{ invoice.amount }}</td>
                    </tr>
                    <tr style="border-bottom: 1px solid #ccc">
                        <td>
                            Online Transaction Fee<sup
                                ><i class="fa fa-info-circle"></i
                            ></sup>
                        </td>
                        <td class="text-right">${{ stripeFee.toFixed(2) }}</td>
                    </tr>
                    <tr>
                        <td>Total</td>
                        <td class="text-right">${{ total.toFixed(2) }}</td>
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
</template>

<script>
export default {
    data() {
        return {
            newInvoiceNumber: "",
            newInvoiceAmount: 0,
            company: "",
            stripeFee: 0,
            invoices: [],
        };
    },
    computed: {
        total() {
            let output = 0;
            this.invoices.map((i) => {
                this.stripeFee = Number(i.amount) * 0.029 + 0.3;
                output += Number(i.amount) + this.stripeFee;
            });
            return output;
        },
        paymentInfo() {
            return {
                amount: this.total,
                invoices: this.invoices,
                company: this.company,
            };
        },
    },
    methods: {
        addInvoice() {
            this.invoices.push({
                number: this.newInvoiceNumber,
                amount: this.newInvoiceAmount,
            });

            this.newInvoiceNumber = "";
            this.newInvoiceAmount = 0;

            this.$refs.invoiceNumber.focus();
        },
    },
};
</script>

<style>
.invoice-card {
    max-width: 500px;
    margin: 0 auto;
    background-color: white;
    border-radius: 5px;
    box-shadow: 3px 3px 30px 0px rgba(0, 0, 0, 0.3);
}
.card-body {
    max-width: 350px;
    width: 100%;
    margin: 0 auto;
    padding: 3rem 1.25rem;
}
.card-header {
    color: white;

    background: #66dc8a;
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
}
table {
    width: 100%;
}
</style>
