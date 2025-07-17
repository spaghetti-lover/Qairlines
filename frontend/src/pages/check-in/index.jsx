"use client";

import { useState, useEffect, useCallback } from "react";
import { useRouter } from "next/router";
import { StepIndicator } from "@/components/checkin/step-indicator";
import { FlightDetailsStep } from "@/components/checkin/flight-details";
import { PassengerListStep } from "@/components/checkin/passenger-list";
import { SeatSelectionStep } from "@/components/checkin/seat-selection";
import { PaymentStep } from "@/components/checkin/payment-step";
import { ConfirmationStep } from "@/components/checkin/confirmation-step";
import LoadingSkeleton from "@/components/checkin/loading-skeleton";
import { useToast } from "@/hooks/use-toast";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;
const API_MAIL_URL = process.env.NEXT_PUBLIC_API_MAIL_URL;

const steps = [
  { title: "Chi tiết chuyến bay", description: "Xem lại thông tin chuyến bay" },
  { title: "Hành khách", description: "Chọn hành khách" },
  { title: "Chọn ghế", description: "Chọn ghế của bạn" },
  { title: "Thanh toán", description: "Thanh toán vé" },
  { title: "Xác nhận", description: "Xác nhận thông tin" },
];

export default function CheckInPage() {
  const [currentStep, setCurrentStep] = useState(0);
  const [bookingID, setBookingID] = useState(null);
  const [email, setEmail] = useState(null);

  const [bookingData, setBookingData] = useState(null);
  const [departureFlight, setDepartureFlight] = useState(null);
  const [returnFlight, setReturnFlight] = useState(null);
  const [passengerList, setPassengerList] = useState({
    departure: [],
    return: [],
  });

  const [departureSeats, setDepartureSeats] = useState([]);
  const [returnSeats, setReturnSeats] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [currentTrip, setCurrentTrip] = useState("departure"); // "departure" hoặc "return"

  const router = useRouter();
  const { toast } = useToast(); // Hook toast

  // ------------------------------------
  // 1. Các hàm generate, tính toán, ...
  // ------------------------------------
  const generateSeatData = useCallback(() => {
    const columns = ["A", "B", "C", "D", "E", "G"];
    const rows = Array.from({ length: 44 }, (_, i) => i + 1);

    return rows.flatMap((row) =>
      columns.map((col) => ({
        id: `${row}${col}`,
        type:
          row === 18 || row === 32
            ? "blocked"
            : Math.random() < 0.3
            ? "unavailable"
            : "available",
      }))
    );
  }, []);

  const calculateDuration = useCallback((departure, arrival) => {
    const durationInMinutes = (arrival - departure) / 60;
    const hours = Math.floor(durationInMinutes / 60);
    const minutes = durationInMinutes % 60;
    return `${hours}h ${minutes}m`;
  }, []);

  const generateGate = useCallback(() => {
    return "6";
  }, []);

  // ------------------------------------
  // 2. Các hàm fetch API
  // ------------------------------------
  const fetchFlightDetails = useCallback(
    async (flightId, type) => {
      try {
        const token = localStorage.getItem("token");
        if (!token) throw new Error("Token không tồn tại.");

        const response = await fetch(`${API_BASE_URL}/api/flight/${flightId}`, {
          headers: { Authorization: `Bearer ${token}` },
        });
        if (!response.ok) {
          throw new Error(`Error fetching flight: ${response.statusText}`);
        }

        const result = await response.json();
        const flightData = result.data;

        const formattedFlight = {
          from: flightData.departureCity,
          to: flightData.arrivalCity,
          departureTime: new Date(
            flightData.departureTime.seconds * 1000
          ).toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
          }),
          arrivalTime: new Date(
            flightData.arrivalTime.seconds * 1000
          ).toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
          }),
          duration: calculateDuration(
            flightData.departureTime.seconds,
            flightData.arrivalTime.seconds
          ),
          flightNumber: flightData.flightNumber,
          date: new Date(
            flightData.departureTime.seconds * 1000
          ).toLocaleDateString("vi-VN"),
        };

        if (type === "departure") {
          setDepartureFlight(formattedFlight);
        } else {
          setReturnFlight(formattedFlight);
        }
      } catch (err) {
        console.error(err);
        setError(`Error fetching ${type} flight: ${err.message}`);
      }
    },
    [calculateDuration, setDepartureFlight, setReturnFlight, setError]
  );

  const fetchTickets = useCallback(
    async (ticketIds) => {
      try {
        const token = localStorage.getItem("token");
        if (!token) throw new Error("Token không tồn tại.");

        const ticketPromises = ticketIds.map(async (ticketId) => {
          const response = await fetch(
            `${API_BASE_URL}/api/ticket?id=${ticketId}`,
            {
              headers: { Authorization: `Bearer ${token}` },
            }
          );
          if (!response.ok) {
            throw new Error(
              `Error fetching ticket ${ticketId}: ${response.statusText}`
            );
          }

          const result = await response.json();
          const ownerData = result.data.owner;
          console.log("Owner Data:", ownerData);
          return {
            id: ticketId,
            title: ownerData.gender === "Female" ? "Bà" : "Ông",
            name: `${ownerData.first_name} ${ownerData.last_name}`,
            type: result.data.flightClass || "Economy",
            flightId: result.data.flightId,
            // Lúc sau bạn có thể thêm thuộc tính "seat" nếu cần
          };
        });

        return await Promise.all(ticketPromises);
      } catch (err) {
        console.error(err);
        setError(err.message);
        return [];
      }
    },
    [setError]
  );

  const fetchBooking = useCallback(async () => {
    if (!bookingID) return;

    try {
      setLoading(true);

      const token = localStorage.getItem("token");
      if (!token) throw new Error("Token không tồn tại.");

      // Fetch booking
      const response = await fetch(
        `${API_BASE_URL}/api/booking?id=${bookingID}`,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      if (!response.ok) {
        throw new Error(`Error fetching booking: ${response.statusText}`);
      }

      const result = await response.json();
      setBookingData(result.data);

      // Fetch flight details
      if (result.data.tripType === "roundTrip") {
        await fetchFlightDetails(result.data.departureFlightId, "departure");
        await fetchFlightDetails(result.data.returnFlightId, "return");
      } else {
        // oneWay
        await fetchFlightDetails(result.data.departureFlightId, "departure");
      }

      // Fetch passengers
      const departurePassengers = await fetchTickets(
        result.data.departureIdTickets
      );
      const returnPassengers =
        result.data.tripType === "roundTrip"
          ? await fetchTickets(result.data.returnIdTickets)
          : [];

      setPassengerList({
        departure: departurePassengers,
        return: returnPassengers,
      });

      // Generate seat data ban đầu
      setDepartureSeats(generateSeatData());
      setReturnSeats(generateSeatData());
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }, [
    bookingID,
    fetchFlightDetails,
    fetchTickets,
    generateSeatData,
    setPassengerList,
    setDepartureSeats,
    setReturnSeats,
    setBookingData,
    setError,
    setLoading,
  ]);

  // ------------------------------------
  // 3. Các useEffect
  // ------------------------------------
  // Lấy bookingID, email từ query
  useEffect(() => {
    if (router.query.bookingID) {
      setBookingID(router.query.bookingID);
    }
    if (router.query.email) {
      setEmail(router.query.email);
    }
  }, [router.query]);

  // Gọi fetchBooking khi bookingID thay đổi
  useEffect(() => {
    fetchBooking();
  }, [fetchBooking]);

  // ------------------------------------
  // 4. Các hàm xử lý sự kiện
  // ------------------------------------
  const handleSeatSelect = (seatId, customerId, tripType) => {
    // Cập nhật seat cho passengerList
    setPassengerList((prev) => ({
      ...prev,
      [tripType]: prev[tripType].map((p) =>
        p.id === customerId ? { ...p, seat: seatId } : p
      ),
    }));

    // Cập nhật trạng thái ghế
    const updateSeats = (seats) =>
      seats.map((seat) =>
        seat.id === seatId
          ? { ...seat, type: "selected" }
          : seat.type === "selected" &&
            passengerList[tripType].some((p) => p.seat === seat.id)
          ? { ...seat, type: "available" }
          : seat
      );

    if (tripType === "departure") {
      setDepartureSeats((prev) => updateSeats(prev));
    } else {
      setReturnSeats((prev) => updateSeats(prev));
    }
  };

  // Hàm gọi API update seat
  const updateSeatsApi = useCallback(async () => {
    try {
      const token = localStorage.getItem("token");
      if (!token) throw new Error("Token không tồn tại.");

      // Lấy seatCode từ passengerList
      const payload = [
        ...passengerList.departure.map((passenger) => ({
          ticketId: passenger.id,
          seatCode: passenger.seat,
        })),
        ...passengerList.return.map((passenger) => ({
          ticketId: passenger.id,
          seatCode: passenger.seat,
        })),
      ].filter((entry) => entry.seatCode);

      if (payload.length === 0) {
        toast({
          title: "Không có ghế",
          description: "Không có ghế nào được chọn để lưu.",
          variant: "destructive",
        });
        return false;
      }

      const response = await fetch(`${API_BASE_URL}/api/ticket/update-seats`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        throw new Error(`API error: ${response.statusText}`);
      }

      await response.json();
      toast({
        title: "Thành công",
        description: "Ghế đã được cập nhật thành công!",
        variant: "default",
      });
      // fetch(`${API_MAIL_URL}/send-mail`, {
      //   method: "POST",
      //   headers: {
      //     "Content-Type": "application/json",
      //   },
      //   body: JSON.stringify({
      //     To: email,
      //     Subject: "Xác nhận ghế máy bay",
      //     Body: `
      //     Xin chào,

      //     Chúng tôi xin thông báo rằng ghế của bạn đã được cập nhật thành công cho chuyến bay.

      //     Mã chuyến bay: ${bookingID}

      //     Vui lòng kiểm tra lại thông tin trong ứng dụng để đảm bảo mọi thứ chính xác.

      //     Chúc bạn có một chuyến bay an toàn và thoải mái!

      //     Trân trọng,
      //     Đội ngũ Qairlines
      //     `,
      //   }),
      // })
      return true;
    } catch (error) {
      console.error("Error updating seats:", error);
      toast({
        title: "Cập nhật thất bại",
        description: "Cập nhật ghế thất bại, vui lòng thử lại.",
        variant: "destructive",
      });
      return false;
    }
  }, [passengerList, toast]);

  // Hàm xử lý điều hướng step
  const handleContinue = async () => {
    // Nếu đang ở step Chọn ghế (index = 2) thì gọi updateSeatsApi
    // if (currentStep === 2) {
    //   const isUpdated = await updateSeatsApi();
    //   if (!isUpdated) return;
    // }
    if (currentStep < steps.length - 1) {
      setCurrentStep((prev) => prev + 1);
    }
  };

  const handleBack = () => {
    if (currentStep > 0) {
      setCurrentStep((prev) => prev - 1);
    }
  };

  // ------------------------------------
  // 5. Render
  // ------------------------------------
  if (loading) {
    return <LoadingSkeleton />;
  }

  if (error) {
    return (
      <div className="container mx-auto p-6 text-red-600">Lỗi: {error}</div>
    );
  }

  return (
    <div className="container mx-auto p-6">
      <StepIndicator currentStep={currentStep} steps={steps} />

      {currentStep === 0 && (
        <FlightDetailsStep
          flightDetails={departureFlight}
          returnFlightDetails={
            bookingData?.tripType === "roundTrip" ? returnFlight : null
          }
          passengerCount={passengerList.departure.length}
          onContinue={handleContinue}
          onCancel={() => window.history.back()}
        />
      )}

      {currentStep === 1 && (
        <PassengerListStep
          passengers={passengerList}
          onContinue={handleContinue}
          onBack={handleBack}
        />
      )}

      {currentStep === 2 && (
        <SeatSelectionStep
          passengers={
            currentTrip === "departure"
              ? passengerList.departure
              : passengerList.return
          }
          seats={currentTrip === "departure" ? departureSeats : returnSeats}
          onSeatSelect={(seatId, customerId) =>
            handleSeatSelect(seatId, customerId, currentTrip)
          }
          onContinue={handleContinue}
          onBack={handleBack}
          onSwitchTrip={() =>
            setCurrentTrip((prev) =>
              prev === "departure" ? "return" : "departure"
            )
          }
          currentTrip={currentTrip}
        />
      )}

      {currentStep === 3 && (
        <PaymentStep
          onPaymentSuccess={async () => {
            const isUpdated = await updateSeatsApi();
            if (isUpdated) setCurrentStep(4);
          }}
          onBack={handleBack}
          bookingId={bookingID}
          amount={10000} // Giả sử 10000 là số tiền thanh toán
          currency="USD"
        />
      )}

      {currentStep === 4 && (
        <ConfirmationStep
          bookingReference={bookingData?.bookingId || "Không rõ"}
          departurePassengers={passengerList.departure || []}
          returnPassengers={passengerList.return || []}
          departureFlight={{
            flightNumber: departureFlight?.flightNumber || "N/A",
            date: departureFlight?.date || "N/A",
            departureTime: departureFlight?.departureTime || "N/A",
            from: departureFlight?.from || "N/A",
            to: departureFlight?.to || "N/A",
            gate: generateGate(),
          }}
          returnFlight={
            bookingData?.tripType === "roundTrip"
              ? {
                  flightNumber: returnFlight?.flightNumber || "N/A",
                  date: returnFlight?.date || "N/A",
                  departureTime: returnFlight?.departureTime || "N/A",
                  from: returnFlight?.from || "N/A",
                  to: returnFlight?.to || "N/A",
                  gate: generateGate(),
                }
              : null
          }
          onBack={handleBack}
          onHome={() => router.push("/")}
        />
      )}
    </div>
  );
}
