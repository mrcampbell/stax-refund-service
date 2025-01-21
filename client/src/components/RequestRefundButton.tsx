import { RequestRefund } from "../services/api-client";

function RequestRefundButton ({paymentID, onRequest}: {paymentID: string, onRequest: Function}) {
    const handleClick = async () => {
        await RequestRefund(paymentID);
        onRequest();
    }

  return (
    <button className="btn btn-info" onClick={handleClick}>Request Refund</button>
  )
}

export default RequestRefundButton;