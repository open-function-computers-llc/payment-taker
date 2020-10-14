<template>
    <div>
        <h1>Which invoices would you like to pay?</h1>
        Invoice Number:<br />
        <input type="text" v-model="newInvoiceNumber" ref="invoiceNumber" autofocus /><br /><br />

        Invoice Amount:<br />
        <input type="number" v-model="newInvoiceAmount" ref="invoiceAmount" /><br />
        <button @click="addInvoice">Add Invoice</button>

        <table v-if="invoices.length" style="width: 400px; margin: 40px auto 0;">
            <tr v-for="(invoice, index) in invoices" :key="'invoice'+index" style="border-bottom: 1px solid #ccc;">
                <td>{{ invoice.number }}</td>
                <td>{{ invoice.amount }}</td>
            </tr>
            <tr>
                <td colspan="2">Total: ${{ total }}</td>
            </tr>
        </table>
        <button v-if="invoices.length" @click="$emit('ready-to-pay', paymentInfo)">Pay</button>
    </div>
</template>

<script>
export default {
    data() {
        return {
            newInvoiceNumber: "",
            newInvoiceAmount: 0,
            invoices: []
        }
    },
    computed: {
        total() {
            let output = 0
            this.invoices.map((i) => {
                output += Number(i.amount)
            })
            return output
        },
        paymentInfo() {
            return {
                amount: this.total,
                invoices: this.invoices
            }
        }
    },
    methods: {
        addInvoice() {
            this.invoices.push({
                number: this.newInvoiceNumber,
                amount: this.newInvoiceAmount
            })

            this.newInvoiceNumber= ""
            this.newInvoiceAmount= 0

            this.$refs.invoiceNumber.focus()
        }
    }
}
</script>
