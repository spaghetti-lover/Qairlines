import React from "react";
import { Plane, Calendar, Clock, MapPin, User, Barcode, Star, DoorOpen } from "lucide-react";
import Image from "next/image";
export default function ModernFlightTicket({
  passengerName,
  flightNumber,
  flightDate,
  departureTime,
  from,
  to,
  flightClass,
  boardingTime,
  gate,
  seat,
  onClose,
}) {
  return (
    <div className="bg-gradient-to-br from-red-400 to-red-600 text-white shadow-xl rounded-xl overflow-hidden w-full max-w-md mx-auto">
      <div className="relative p-6">
        {/* Header */}
        <div className="flex justify-between items-center mb-6">
        <Image 
          src="/logo-white.png" 
          alt="Logo" 
          className="w-36" 
          width={144} 
          height={144} 
        />
          <h2 className="text-2xl font-bold">QAirlines</h2>
        </div>

        {/* Ticket Details */}
        <div className="space-y-6">
          {/* Passenger and Flight */}
          <div className="flex justify-between items-center">
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Hành Khách</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <User className="w-5 h-5 mr-2" />
                {passengerName}
              </p>
            </div>
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Chuyến bay</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <Plane className="w-5 h-5 mr-2" />
                <p className="text-lg font-semibold mt-1">{flightNumber}</p>
              </p>  
            </div>
          </div>

          {/* Date and Departure */}
          <div className="flex justify-between items-center">
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Ngày</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <Calendar className="w-5 h-5 mr-2" />
                {flightDate}
              </p>
            </div>
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Giờ khởi hành</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <Clock className="w-5 h-5 mr-2" />
                {departureTime}
              </p>
            </div>
          </div>

          {/* From and To */}
          <div className="flex justify-between items-center">
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Từ</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <MapPin className="w-5 h-5 mr-2" />
                {from}
              </p>
            </div>
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Đến</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <MapPin className="w-5 h-5 mr-2" />
                {to}
              </p>
            </div>
          </div>

          {/* Class and Boarding */}
          <div className="flex justify-between items-center">
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Hạng vé</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <Star className="w-5 h-5 mr-2" />
                {flightClass}
              </p>
            </div>
            <div>
              <p className="text-xs uppercase tracking-wide opacity-75">Giờ lên cổng</p>
              <p className="text-lg font-semibold flex items-center mt-1">
                <DoorOpen className="w-5 h-5 mr-2" />
                {boardingTime}
              </p>
            </div>
          </div>
        </div>
      </div>

      {/* Footer */}
      <div className="bg-white text-zinc-800 p-6 flex justify-between items-center">
        <div>
          <p className="text-xs uppercase tracking-wide text-zinc-500">Cổng</p>
          <p className="text-2xl font-bold mt-1">{gate}</p>
        </div>
        <div>
          <p className="text-xs uppercase tracking-wide text-zinc-500">Chỗ Ngồi</p>
          <p className="text-2xl font-bold mt-1">{seat}</p>
        </div>
        <Barcode className="w-24 h-24 text-zinc-800" />
      </div>
    </div>
  );
}
