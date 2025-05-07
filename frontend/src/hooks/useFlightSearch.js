import { useState, useEffect, useCallback } from 'react';

const useFlightSearch = () => {
  const [fromAirport, setFromAirport] = useState('');
  const [toAirport, setToAirport] = useState('');
  const [departureDate, setDepartureDate] = useState(null);
  const [returnDate, setReturnDate] = useState(null);
  const [passengerCount, setPassengerCount] = useState(1);
  const [tripType, setTripType] = useState('roundTrip');
  const [errors, setErrors] = useState({});
  const [isValid, setIsValid] = useState(false);

  const swapAirports = () => {
    setFromAirport(toAirport);
    setToAirport(fromAirport);
  };

  // Bọc validateInputs trong useCallback
  const validateInputs = useCallback(() => {
    const newErrors = {};

    if (!fromAirport) newErrors.fromAirport = 'Vui lòng chọn sân bay đi';
    if (!toAirport) newErrors.toAirport = 'Vui lòng chọn sân bay đến';
    if (!departureDate) newErrors.departureDate = 'Vui lòng chọn ngày đi';
    if (tripType === 'roundTrip' && !returnDate)
      newErrors.returnDate = 'Vui lòng chọn ngày về';

    if (!Number.isInteger(passengerCount) || passengerCount < 1) {
      newErrors.passengerCount = 'Số lượng hành khách không hợp lệ';
    }

    if (departureDate && returnDate) {
      if (new Date(returnDate) <= new Date(departureDate)) {
        newErrors.returnDate = 'Ngày về phải sau ngày đi';
      }
    }

    setErrors(newErrors);
    setIsValid(Object.keys(newErrors).length === 0);
  }, [
    fromAirport,
    toAirport,
    departureDate,
    returnDate,
    passengerCount,
    tripType,
    setErrors,
    setIsValid,
  ]);

  // Gọi hàm validateInputs bên trong useEffect
  useEffect(() => {
    validateInputs();
  }, [validateInputs]);

  return {
    fromAirport,
    setFromAirport,
    toAirport,
    setToAirport,
    departureDate,
    setDepartureDate,
    returnDate,
    setReturnDate,
    passengerCount,
    setPassengerCount,
    tripType,
    setTripType,
    swapAirports,
    errors,
    isValid,
  };
};

export default useFlightSearch;
