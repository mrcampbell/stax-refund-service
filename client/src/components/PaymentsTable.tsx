import RequestRefundButton from "./RequestRefundButton";

function PaymentsTable({payments, refunds, onUpdate}: {payments: any[], refunds: any[], onUpdate: Function}) {

    const merged = payments.map((payment) => {
        const refund = refunds.find((refund) => refund.payment_id === payment.id);
        return {
            ...payment,
            refund: refund
        }});

        console.log(merged)
    return (
        <div className="overflow-x-auto">
            <table className="table table-zebra">
                {/* head */}
                <thead>
                    <tr>
                        <th></th>
                        <th>ID</th>
                        <th>Description</th>
                        <th>Amount</th>
                        <th>Timestamp</th>
                        <th>Refund</th>
                    </tr>
                </thead>
                <tbody>
                    {merged.map((payment) => (
                        <tr key={payment.id}>
                            <td></td>
                            <td>{payment.id}</td>
                            <td>{payment.description}</td>
                            <td>{payment.amount}</td>
                            <td>{payment.timestamp}</td>
                            <td>{payment.refund ? refundStatusToElement(payment.refund.status) : <RequestRefundButton paymentID={payment.id} onRequest={onUpdate} />}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>)
}

function refundStatusToElement(status: number) {
    switch (status) {
        case 0:
            return <div className="badge badge-neutral">neutral</div>
        case 1:
            return <div className="badge badge-warning">pending</div>
        case 2:
            return <div className="badge badge-success">approved</div>
        case 3:
            return <div className="badge badge-danger">failure</div>
    }


}

export default PaymentsTable