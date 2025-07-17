import React, { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { loadStripe } from "@stripe/stripe-js";
import {
  Elements,
  PaymentElement,
  useStripe,
  useElements,
} from "@stripe/react-stripe-js";

const stripePromise = loadStripe(
  process.env.NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY
);

function StripeForm({ onPaymentSuccess }) {
  const stripe = useStripe();
  const elements = useElements();
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage("");
    if (!stripe || !elements) {
      setMessage("Stripe chưa sẵn sàng.");
      setLoading(false);
      return;
    }
    const { error } = await stripe.confirmPayment({
      elements,
      confirmParams: {},
      redirect: "if_required",
    });

    if (error) {
      setMessage(error.message || "Thanh toán thất bại.");
      setLoading(false);
      return;
    }

    setMessage("Thanh toán thành công!");
    setLoading(false);
    if (onPaymentSuccess) onPaymentSuccess();
  };

  return (
    <form
      id="payment-form"
      onSubmit={handleSubmit}
      className="w-full max-w-md mx-auto p-6 bg-white rounded shadow"
    >
      <div id="payment-element">
        <PaymentElement />
      </div>
      <div id="error-message" className="mt-2 text-red-600 text-sm">
        {message}
      </div>
      <Button
        type="submit"
        disabled={loading}
        variant="orange"
        className="mt-6 w-full"
      >
        {loading ? "Đang xử lý..." : "Pay"}
      </Button>
    </form>
  );
}

export function PaymentStep({
  onPaymentSuccess,
  onBack,
  bookingId,
  amount,
  currency,
}) {
  const [ready, setReady] = useState(false);
  const [clientSecret, setClientSecret] = useState("");

  useEffect(() => {
    // Chỉ gọi khi có đủ thông tin
    if (bookingId && amount && currency) {
      fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/payment-intents`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          booking_id: parseInt(bookingId),
          amount: parseInt(amount),
          currency: currency,
        }),
      })
        .then((res) => res.json())
        .then((data) => {
          setClientSecret(data.client_secret);
          setReady(true);
        })
        .catch(() => setReady(false));
    }
  }, [bookingId, amount, currency]);

  const options = {
    clientSecret,
    appearance: {
      theme: "stripe",
    },
  };
  return (
    <div className="sr-root flex flex-col items-center justify-center min-h-[400px] bg-gray-50">
      <div className="sr-main w-full max-w-md">
        <h1 className="text-2xl font-bold mb-6">Thanh toán</h1>
        <div className="mb-4 text-lg font-semibold text-orange-700 text-center">
          Số tiền cần thanh toán: {amount} {currency?.toUpperCase()}
        </div>
        {ready ? (
          <Elements stripe={stripePromise} options={options}>
            <StripeForm onPaymentSuccess={onPaymentSuccess} />
          </Elements>
        ) : (
          <div className="text-center text-gray-500">
            Đang tải giao diện thanh toán...
          </div>
        )}
        <Button onClick={onBack} variant="outline" className="mt-4 w-full">
          Quay lại
        </Button>
      </div>
    </div>
  );
}
