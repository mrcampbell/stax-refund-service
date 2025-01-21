import { useEffect, useState } from "react";
import { GetPayments, GetRefunds } from "../services/api-client";
import PaymentsTable from "../components/PaymentsTable";

function RefundPage() {
    const [payments, setPayments] = useState([]);
    const [refunds, setRefunds] = useState([]);

    const fetchPayments = async () => {
        const payments = await GetPayments()
        setPayments(payments);
        console.log(payments);
    }
    
    const fetchRefunds = async () => {
        const refunds = await GetRefunds();
        setRefunds(refunds);
        console.log(refunds)
    }

    useEffect(() => {
        fetchPayments();
        fetchRefunds();
    }, []);

    const handleUpdate = () => {
        fetchPayments();
        fetchRefunds();
    }

    return (
        <div className="p-8">
            <h1>Refund Page</h1>
            <PaymentsTable payments={payments} refunds={refunds} onUpdate={handleUpdate} />
        </div>
    );
}

export default RefundPage;