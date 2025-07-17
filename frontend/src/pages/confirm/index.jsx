import React from "react";
import { motion } from 'framer-motion'
import { Button } from "@/components/ui/button";
import { CheckCircle, Calendar, Clock, Plane, CreditCard, Users } from 'lucide-react';
import { Card, CardContent } from "@/components/ui/card";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from "@/components/ui/dialog";
import { PassengerInfoDialog } from "@/components/PassengerInfoDialog";
import { Skeleton } from "@/components/ui/skeleton";
import { useFlightConfirmation } from "@/hooks/useFlightConfirmation";

export default function ConfirmationPage() {
  const {
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
  } = useFlightConfirmation();

  const formatTime = (seconds) => {
    return new Date(seconds * 1000).toLocaleTimeString("vi-VN", {
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  const calculateFlightDuration = (departureTime, arrivalTime) => {
    const durationInMinutes = (arrivalTime - departureTime) / 60;
    const hours = Math.floor(durationInMinutes / 60);
    const minutes = Math.floor(durationInMinutes % 60);
    return `${hours} giờ ${minutes} phút`;
  };

  if (loading) {
    return <LoadingSkeleton />;
  }

  if (error) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <p className="text-xl text-gray-600">{error}</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <HeroSection />
      <FlightDetails
        tripType={tripType}
        departureFlightData={departureFlightData}
        returnFlightData={returnFlightData}
        departureOption={departureOption}
        returnOption={returnOption}
        formatTime={formatTime}
        calculateFlightDuration={calculateFlightDuration}
        passengerCount={passengerCount}
      />
      <TotalAndActions
        totalAmount={totalAmount}
        isPassengerInfoFilled={isPassengerInfoFilled}
        handleOpenPassengerInfo={handleOpenPassengerInfo}
        handleConfirmPayment={handleConfirmPayment}
      />
      <ConfirmationDialog
        isOpen={isPaymentConfirmed}
        onOpenChange={setIsPaymentConfirmed}
        bookingId={bookingId}
        handleReturnHome={handleReturnHome}
      />
      <PassengerInfoDialog
        isOpen={isPassengerInfoOpen}
        onClose={() => setIsPassengerInfoOpen(false)}
        passengerCount={parseInt(passengerCount) || 1}
        onInfoFilled={(info) => {
          handlePassengerInfoFilled();
          handleSavePassengerInfo(info);
          setIsPassengerInfoOpen(false);
        }}
      />
    </div>
  );
}

function LoadingSkeleton() {
  return (
    <div className="min-h-screen bg-gray-50 p-4 sm:p-6">
      <div className="h-[200px] w-full bg-orange relative mb-4">
        <div className="absolute inset-0 flex flex-col items-center justify-center text-white px-4 text-center">
          <Skeleton className="h-8 w-1/2 mb-2" />
          <Skeleton className="h-6 w-2/3" />
        </div>
      </div>

      <div className="max-w-4xl mx-auto">
        <Card className="shadow-lg border-orange">
          <CardContent className="p-4 sm:p-6">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
              <Skeleton className="h-6 w-1/3 mb-2 sm:mb-0" />
              <Skeleton className="h-4 w-20" />
            </div>
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-8">
              <div className="flex items-center space-x-4 mb-4 sm:mb-0">
                <Skeleton className="h-8 w-16" />
                <div className="flex-1 relative px-8">
                  <Skeleton className="h-6 w-6 mx-auto" />
                </div>
                <Skeleton className="h-8 w-16" />
              </div>
              <div className="text-right">
                <Skeleton className="h-4 w-32 mb-1" />
                <Skeleton className="h-6 w-24" />
              </div>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6 mb-8">
              <div className="flex items-center space-x-3">
                <Skeleton className="h-5 w-5 rounded-full" />
                <div>
                  <Skeleton className="h-4 w-24 mb-1" />
                  <Skeleton className="h-4 w-16" />
                </div>
              </div>
              <div className="flex items-center space-x-3">
                <Skeleton className="h-5 w-5 rounded-full" />
                <div>
                  <Skeleton className="h-4 w-24 mb-1" />
                  <Skeleton className="h-4 w-16" />
                </div>
              </div>
              <div className="flex items-center space-x-3">
                <Skeleton className="h-5 w-5 rounded-full" />
                <div>
                  <Skeleton className="h-4 w-24 mb-1" />
                  <Skeleton className="h-4 w-16" />
                </div>
              </div>
              <div className="flex items-center space-x-3">
                <Skeleton className="h-5 w-5 rounded-full" />
                <div>
                  <Skeleton className="h-4 w-24 mb-1" />
                  <Skeleton className="h-4 w-16" />
                </div>
              </div>
            </div>

            <div className="border-t border-gray-200 pt-6">
              <div className="flex flex-col sm:flex-row sm:justify-between sm:items-center mb-4">
                <div className="mb-2 sm:mb-0">
                  <Skeleton className="h-4 w-20 mb-1" />
                  <Skeleton className="h-3 w-32" />
                </div>
                <Skeleton className="h-6 w-32" />
              </div>
              <div className="flex space-x-4">
                <Skeleton className="h-10 w-48 rounded" />
                <Skeleton className="h-10 w-48 rounded" />
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}

function HeroSection() {
  return (
    <div className="relative h-[200px] w-full bg-orange">
      <div className="absolute inset-0 flex flex-col items-center justify-center text-white px-4 text-center">
        <h1 className="text-2xl md:text-3xl font-bold mb-2">Lựa chọn của quý hành khách</h1>
      </div>
    </div>
  );
}

function FlightDetails({ tripType, departureFlightData, returnFlightData, departureOption, returnOption, formatTime, calculateFlightDuration, passengerCount }) {
  const flightCards = [
    { flightData: departureFlightData, option: departureOption, title: "Chi tiết chuyến bay đi" },
    ...(tripType === "roundTrip" ? [{ flightData: returnFlightData, option: returnOption, title: "Chi tiết chuyến bay về" }] : []),
  ].filter(Boolean);

  return (
    <div className="max-w-4xl mx-auto p-4 sm:p-6 m-4 relative z-10">
      {flightCards.map((card, index) => (
        <FlightCard
          key={index}
          {...card}
          formatTime={formatTime}
          calculateFlightDuration={calculateFlightDuration}
          passengerCount={passengerCount}
        />
      ))}
    </div>
  );
}

function FlightCard({ flightData, option, title, formatTime, calculateFlightDuration, passengerCount }) {
  return (
    <Card className="shadow-lg border-orange mb-4">
      <CardContent className="p-4 sm:p-6">
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
          <h2 className="text-xl sm:text-2xl font-semibold text-gray-800 mb-2 sm:mb-0">{title}</h2>
          <span className="text-sm text-gray-500">{flightData.flightNumber}</span>
        </div>

        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-8">
          <div className="flex items-center space-x-4 mb-4 sm:mb-0">
            <div className="flex flex-col items-center">
              <span className="text-2xl sm:text-3xl font-bold text-gray-800">
                {formatTime(flightData.departureTime.seconds)}
              </span>
              <span className="text-base sm:text-lg font-medium text-gray-600">
                {flightData.departureCity}
              </span>
            </div>
            <div className="flex-1 relative px-8">
              <Plane className="text-orange absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 transform rotate-45" />
            </div>
            <div className="flex flex-col items-center">
              <span className="text-2xl sm:text-3xl font-bold text-gray-800">
                {formatTime(flightData.arrivalTime.seconds)}
              </span>
              <span className="text-base sm:text-lg font-medium text-gray-600">
                {flightData.arrivalCity}
              </span>
            </div>
          </div>
          <div className="text-right">
            <div className="text-sm font-medium text-gray-500">{option.name}</div>
            <div className="text-xl sm:text-2xl font-bold text-orange">{option.price.toLocaleString()} VND</div>
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6 mb-8">
          <div className="flex items-center space-x-3">
            <Calendar className="text-orange flex-shrink-0" />
            <div>
              <div className="text-sm text-gray-500">Ngày khởi hành</div>
              <div className="font-medium">
                {new Date(flightData.departureTime.seconds * 1000).toLocaleDateString("vi-VN")}
              </div>
            </div>
          </div>
          <div className="flex items-center space-x-3">
            <Clock className="text-orange flex-shrink-0" />
            <div>
              <div className="text-sm text-gray-500">Thời gian bay</div>
              <div className="font-medium">
                {calculateFlightDuration(flightData.departureTime.seconds, flightData.arrivalTime.seconds)}
              </div>
            </div>
          </div>
          <div className="flex items-center space-x-3">
            <Users className="text-orange flex-shrink-0" />
            <div>
              <div className="text-sm text-gray-500">Hành khách</div>
              <div className="font-medium">{passengerCount} Người lớn/ Trẻ em</div>
            </div>
          </div>
          <div className="flex items-center space-x-3">
            <CreditCard className="text-orange flex-shrink-0" />
            <div>
              <div className="text-sm text-gray-500">Phương thức thanh toán</div>
              <div className="font-medium">Thẻ tín dụng</div>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}

function TotalAndActions({ totalAmount, isPassengerInfoFilled, handleOpenPassengerInfo, handleConfirmPayment }) {
  return (
    <div className="border-t border-gray-200 pt-6 pb-8">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col sm:flex-row sm:justify-between sm:items-center mb-4">
          <div className="mb-2 sm:mb-0">
            <div className="text-lg font-semibold text-gray-800">Tổng cộng</div>
            <div className="text-sm text-gray-500">Đã bao gồm thuế và phí</div>
          </div>
          <div className="text-2xl sm:text-3xl font-bold text-orange">
            {totalAmount.toLocaleString()} VND
          </div>
        </div>
        <div className="flex flex-col sm:flex-row space-y-3 sm:space-y-0 sm:space-x-3">
          <Button
            variant="outline"
            className="w-full sm:w-auto"
            onClick={handleOpenPassengerInfo}
          >
            Nhập thông tin hành khách
          </Button>
          <Button
            variant="orange"
            className={`w-full sm:w-auto text-white ${!isPassengerInfoFilled && "opacity-50 cursor-not-allowed"}`}
            onClick={handleConfirmPayment}
            disabled={!isPassengerInfoFilled}
          >
            Xác nhận và thanh toán
          </Button>
        </div>
      </div>
    </div>
  );
}

function ConfirmationDialog({ isOpen, onOpenChange, bookingId, handleReturnHome }) {
  return (
    <Dialog open={isOpen} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <motion.div
            initial={{ scale: 0 }}
            animate={{ scale: 1 }}
            transition={{ type: "spring", stiffness: 260, damping: 20 }}
            className="flex justify-center mb-4"
          >
            <CheckCircle className="w-16 h-16 text-green-500" />
          </motion.div>
          <DialogTitle className="text-2xl font-bold text-center">Thanh toán thành công</DialogTitle>
        </DialogHeader>
        <div className="flex flex-col items-center space-y-4 text-center">
          <DialogDescription className="text-lg">
            Mã đặt vé của bạn là:
          </DialogDescription>
          <div className="px-4 py-2 text-xl font-mono font-bold bg-gray-100 rounded-md">
            {bookingId}
          </div>
          <DialogDescription className="text-base">
            Cảm ơn quý khách đã đặt vé. Chúc quý khách có chuyến bay vui vẻ!
          </DialogDescription>
        </div>
        <Button
          onClick={handleReturnHome}
          variant="orange"
          className="w-full mt-6 text-white transition-colors duration-200"
        >
          Quay về trang chủ
        </Button>
      </DialogContent>
    </Dialog>
  );
}

