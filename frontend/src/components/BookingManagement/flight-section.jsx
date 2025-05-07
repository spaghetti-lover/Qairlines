'use client';

import { useState } from 'react';
import { ChevronDown, ChevronUp } from 'lucide-react';
import { FlightInfoCard } from './flight-info-card';
import { FlightDetailCard } from './flight-detail-card';
import { Button } from '@/components/ui/button';
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible';

export function FlightSection({
  type,
  flightNumber,
  departureTime,
  arrivalTime,
  departureCity,
  arrivalCity,
  date,
  duration,
  passengers,
  paymentMethod,
  passengerDetails,
  ticketRef,
}) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="mb-8">
      <FlightDetailCard
        type={type}
        flightNumber={flightNumber}
        departureTime={departureTime}
        arrivalTime={arrivalTime}
        departureCity={departureCity}
        arrivalCity={arrivalCity}
        date={date}
        duration={duration}
        passengers={passengers}
        paymentMethod={paymentMethod}
      />
      
      <Collapsible open={isOpen} onOpenChange={setIsOpen}>
        <CollapsibleTrigger asChild>
          <Button
            variant="outline"
            className="w-full flex items-center justify-between py-2 px-4 mt-2 text-primary hover:text-primary-foreground hover:bg-orange transition-colors"
          >
            <span className="text-sm font-medium">
              {isOpen ? 'Ẩn thông tin vé' : 'Hiện thông tin vé'}
            </span>
            {isOpen ? <ChevronUp className="h-4 w-4" /> : <ChevronDown className="h-4 w-4" />}
          </Button>
        </CollapsibleTrigger>
        <CollapsibleContent className="transition-all duration-300 ease-in-out">
          <div className="space-y-4 pt-1">
            {passengerDetails.map((detail, index) => (
              <FlightInfoCard
                key={index}
                type={type}
                passengerName={`${detail.firstName} ${detail.lastName}`}
                seatNumber={detail.seatCode || 'Chưa xác định'}
                onView={detail.onView} 
                onDownload={detail.onDownload} // Sử dụng hàm đã truyền
                ticketRef={ticketRef}  // Truyền hàm onDownload vào FlightInfoCard
              />
            ))}
          </div>
        </CollapsibleContent>
      </Collapsible>
    </div>
  );
}
