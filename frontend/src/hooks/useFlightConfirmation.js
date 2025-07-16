import { useState, useEffect, useCallback } from "react";
import { useRouter } from "next/router";
import { format, parse } from "date-fns";
import { useAccountInfo } from "@/hooks/useAccountInfo";
import { toast } from "@/hooks/use-toast";

/** Khai báo constant để ngoài cùng, đảm bảo không thay đổi reference */
const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

/**
 * Hàm tiện ích tách ra ngoài để không bị re-create mỗi lần render.
 * Không phụ thuộc state/props => có thể định nghĩa ngoài hook
 */
function generateTicketOptions(basePrice, type) {
  const changeFee = type === "economy" ? 860000 : 360000;
  const refundFee = type === "economy" ? 860000 : 360000;
  const checkedBaggage = type === "economy" ? "1 x 23 kg" : "2 x 32 kg";
  const carryOn = "Không quá 12kg";
  return [
    {
      id: `${type}1`,
      name: type === "economy" ? "Phổ Thông Tiêu Chuẩn" : "Thương Gia Tiêu Chuẩn",
      price: basePrice,
      changeFee,
      refundFee,
      checkedBaggage,
      carryOn,
    },
    {
      id: `${type}2`,
      name: type === "economy" ? "Phổ Thông Linh Hoạt" : "Thương Gia Linh Hoạt",
      price: basePrice + 500000,
      changeFee: changeFee / 2,
      refundFee: refundFee / 2,
      checkedBaggage,
      carryOn,
    },
  ];
}

export function useFlightConfirmation() {
  const router = useRouter();
  const {
    departureFlightId,
    departureOptionId,
    returnFlightId,
    returnOptionId,
    passengerCount,
  } = router.query;

  const tripType = returnFlightId && returnOptionId ? "roundTrip" : "oneWay";

  const [departureFlightData, setDepartureFlightData] = useState(null);
  const [returnFlightData, setReturnFlightData] = useState(null);
  const [departureOption, setDepartureOption] = useState(null);
  const [returnOption, setReturnOption] = useState(null);
  const [isPaymentConfirmed, setIsPaymentConfirmed] = useState(false);
  const [isPassengerInfoFilled, setIsPassengerInfoFilled] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isPassengerInfoOpen, setIsPassengerInfoOpen] = useState(false);
  const [bookingId, setBookingId] = useState(null);

  const { personalInfo, loading: accountLoading } = useAccountInfo();

  /**
   * fetchFlightData cần đưa vào useCallback để "ổn định" reference.
   * Bổ sung [setError] vào dependency (nếu có thêm phụ thuộc khác, thêm tiếp).
   */
  const fetchFlightData = useCallback(
    async (flightId, optionId, setFlightDataFn, setOptionFn) => {
      try {
        const response = await fetch(`${API_BASE_URL}/api/flight/${flightId}`);
        if (!response.ok) {
          throw new Error("Không thể lấy dữ liệu chuyến bay.");
        }
        const result = await response.json();
        const data = result.data;

        const { departureCity, arrivalCity, flightId: flightDataId } = data;

        // Gọi hàm tiện ích đã tách ra ngoài
        const economyOptions = generateTicketOptions(data.basePrice, "economy");
        const businessOptions = generateTicketOptions(data.basePrice * 1.5, "business");
        const allOptions = [...economyOptions, ...businessOptions];

        const option = allOptions.find((opt) => opt.id === optionId);
        if (!option) {
          throw new Error("Không tìm thấy thông tin hạng vé.");
        }

        setFlightDataFn({ ...data, departureCity, arrivalCity, flightDataId });
        setOptionFn(option);
      } catch (err) {
        setError(err.message);
      }
    },
    [setError]
  );

  /**
   * useEffect: gọi fetch khi đã có departureFlightId, departureOptionId.
   * => Thêm fetchFlightData vào dependencies để ESLint không cảnh báo.
   */
  useEffect(() => {
    if (!departureFlightId || !departureOptionId) return;

    setLoading(true);
    const fetchData = async () => {
      await fetchFlightData(
        departureFlightId,
        departureOptionId,
        setDepartureFlightData,
        setDepartureOption
      );
      if (returnFlightId && returnOptionId) {
        await fetchFlightData(
          returnFlightId,
          returnOptionId,
          setReturnFlightData,
          setReturnOption
        );
      }
      setLoading(false);
    };

    fetchData();
  }, [
    departureFlightId,
    departureOptionId,
    returnFlightId,
    returnOptionId,
    fetchFlightData, // Quan trọng: thêm fetchFlightData ở đây
  ]);

  const totalAmount =
    (departureOption?.price + (returnOption?.price || 0)) *
    parseInt(passengerCount || 1, 10);

  const handlePassengerInfoFilled = () => {
    setIsPassengerInfoFilled(true);
  };

  const handleConfirmPayment = () => {
    setIsPaymentConfirmed(true);
    if (!isPassengerInfoFilled) {
      toast({
        title: "Thông tin chưa đầy đủ",
        description: "Vui lòng nhập thông tin hành khách trước khi thanh toán.",
        variant: "destructive",
      });
      return;
    }

    toast({
      title: "Thanh toán thành công",
      description: "Cảm ơn quý khách đã đặt vé. Chúc quý khách có chuyến bay vui vẻ!",
    });
  };

  const handleReturnHome = () => {
    router.push("/");
  };

  const handleOpenPassengerInfo = () => {
    setIsPassengerInfoOpen(true);
  };

  const handleSavePassengerInfo = async (passengerData) => {
    const departureTicketDataList = passengerData.map((info) => ({
      price: departureOption.price,
      flightClass: departureOption.name.includes("Thương Gia") ? "business" : "economy",
      ownerData: {
        identityCardNumber: info.idNumber,
        firstName: info.firstName,
        lastName: info.lastName,
        phoneNumber: info.phoneNumber,
        dateOfBirth: format(parse(info.birthDate, "dd/MM/yyyy", new Date()), "yyyy-MM-dd"),
        gender: info.gender,
        address: info.address,
      },
    }));

    const returnTicketDataList = passengerData.map((info) => ({
      price: returnOption?.price || 0,
      flightClass: returnOption?.name.includes("Thương Gia") ? "business" : "economy",
      ownerData: {
        identityCardNumber: info.idNumber,
        firstName: info.firstName,
        lastName: info.lastName,
        phoneNumber: info.phoneNumber,
        dateOfBirth: format(parse(info.birthDate, "dd/MM/yyyy", new Date()), "yyyy-MM-dd"),
        gender: info.gender,
        address: info.address,
      },
    }));

    const bookingData = {
      // bookerId: personalInfo?.uid,
      departureCity: departureFlightData?.departureCity,
      arrivalCity: departureFlightData?.arrivalCity,
      departureFlightId: departureFlightData?.flightId,
      returnFlightId: tripType === "roundTrip" ? returnFlightData?.flightId : undefined,
      tripType,
      departureTicketDataList,
      returnTicketDataList: tripType === "roundTrip" ? returnTicketDataList : undefined,
    };

    try {
      const response = await fetch(`${API_BASE_URL}/api/booking`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify(bookingData),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Lỗi khi tạo booking.");
      }
      const result = await response.json();
      setBookingId(result.data.bookingId);
      toast({
        title: "Đặt vé thành công",
        description: `Mã đặt vé của bạn là: ${result.data.bookingId}`,
        variant: "success",
      });
    } catch (error) {
      console.log(JSON.stringify(bookingData))
      toast({
        title: "Lỗi đặt vé",
        description: error.message,
        variant: "destructive",
      });
    }
  };

  return {
    tripType,
    departureFlightData,
    returnFlightData,
    departureOption,
    returnOption,
    isPaymentConfirmed,
    isPassengerInfoFilled,
    loading,
    error,
    isPassengerInfoOpen,
    bookingId,
    totalAmount,
    passengerCount,
    handlePassengerInfoFilled,
    handleConfirmPayment,
    handleReturnHome,
    handleOpenPassengerInfo,
    handleSavePassengerInfo,
    setIsPaymentConfirmed,
    setIsPassengerInfoOpen,
  };
}
