import { SetLocalToken } from "./auth-service";

export const Authenticate = async (username: string, password: string) => {
    const response = await fetch('http://localhost:8080/api/users/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
    });

    const result = await response.json();
    const { token } = result;
    SetLocalToken(token);
    return { token };
}

export const GetPayments = async () => {
  return authedGet('/api/payments');
}

export const GetRefunds = async () => {
    return authedGet('/api/refunds');
}

export const RequestRefund = async (paymentID: string) => {
    return authedPost('/api/refunds', { payment_id: paymentID, description: `Refund requested on ${new Date().getDate()}` });
}

const authedGet = async (endpoint: string) => {
    const token = localStorage.getItem('token');
    const response = await fetch(`http://localhost:8080${endpoint}`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        }
    });

    const result = await response.json();
    return result;
}

const authedPost = async (endpoint: string, body: any) => {
    const token = localStorage.getItem('token');
    const response = await fetch(`http://localhost:8080${endpoint}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(body)
    });

    const result = await response.json();
    return result;
}