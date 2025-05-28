import { useRouter } from 'next/router';
import { Download } from 'lucide-react';
import { FlightSection } from '@/components/BookingManagement/flight-section';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogOverlay } from '@/components/ui/dialog';
import ModernFlightTicket from '@/components/checkin/flight-ticket';
import { useFlightBooking } from '@/hooks/useFlightBooking';
import { Skeleton } from '@/components/ui/skeleton';
import { useToast } from '@/hooks/use-toast';

export default function FlightBookingPage() {
  const router = useRouter();
  const { bookingID } = router.query;
  const { toast } = useToast();

  const {
    bookingData,
    departureTicketData,
    returnTicketData,
    departureFlightData,
    returnFlightData,
    selectedTicket,
    dialogOpen,
    setDialogOpen,
    error,
    ticketRef,
    handleViewTicket,
    handleDownload,
    handleCancelTicket,
  } = useFlightBooking(bookingID, {
    // Giả sử ta bổ sung option gọi callback khi hủy vé thành công hoặc lỗi
    onCancelSuccess: () => {
      toast({
        title: "Hủy vé thành công",
        description: "Vé của bạn đã được hủy.",
        variant: "default",
      });
    },
    onCancelError: (errMsg) => {
      toast({
        title: "Hủy vé thất bại",
        description: errMsg || "Đã xảy ra lỗi, vui lòng thử lại.",
        variant: "destructive",
      });
    }
  });

  if (error) {
    toast({
      title: "Lỗi",
      description: error,
      variant: "destructive",
    });
    return (
      <div className="container mx-auto py-8 max-w-2xl">
        <h1 className="text-3xl font-bold text-primary mb-6">Lỗi</h1>
        <p>{error}</p>
      </div>
    );
  }

  const isLoading = !bookingData || !departureFlightData || (bookingData.returnFlightId && !returnFlightData);

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 max-w-2xl">
        <div className="space-y-4">
          <Skeleton className="h-8 w-full" />
          <Skeleton className="h-8 w-full" />
          <Skeleton className="h-8 w-3/4" />
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto py-8 max-w-2xl">
      <h1 className="text-3xl font-bold text-primary mb-6">Chi tiết chuyến bay</h1>
      <FlightSection
        type="departure"
        flightNumber={departureFlightData.flightNumber}
        departureTime={new Date(departureFlightData.departureTime.seconds * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
        arrivalTime={new Date(departureFlightData.arrivalTime.seconds * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
        departureCity={departureFlightData.departureCity}
        arrivalCity={departureFlightData.arrivalCity}
        date={new Date(departureFlightData.departureTime.seconds * 1000).toLocaleDateString('vi-VN', { day: '2-digit', month: '2-digit', year: 'numeric' })}
        duration={`${Math.round((departureFlightData.arrivalTime.seconds - departureFlightData.departureTime.seconds) / 60)} phút`}
        passengers={departureTicketData.length}
        paymentMethod="Thẻ tín dụng"
        passengerDetails={departureTicketData.map((ticket) => ({
          firstName: ticket.owner.firstName,
          lastName: ticket.owner.lastName,
          seatCode: ticket.seatCode,
          flightClass: ticket.flightClass,
          onView: () => handleViewTicket(ticket),
          onDownload: () => handleDownload(ticketRef.current),
        }))}
        ticketRef={ticketRef}
      />
      {returnFlightData && (
        <FlightSection
          type="return"
          flightNumber={returnFlightData.flightNumber}
          departureTime={new Date(returnFlightData.departureTime.seconds * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
          arrivalTime={new Date(returnFlightData.arrivalTime.seconds * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
          departureCity={returnFlightData.departureCity}
          arrivalCity={returnFlightData.arrivalCity}
          date={new Date(returnFlightData.departureTime.seconds * 1000).toLocaleDateString('vi-VN', { day: '2-digit', month: '2-digit', year: 'numeric' })}
          duration={`${Math.round((returnFlightData.arrivalTime.seconds - returnFlightData.departureTime.seconds) / 60)} phút`}
          passengers={returnTicketData.length}
          paymentMethod="Thẻ tín dụng"
          passengerDetails={returnTicketData.map((ticket) => ({
            firstName: ticket.owner.firstName,
            lastName: ticket.owner.lastName,
            seatCode: ticket.seatCode,
            flightClass: ticket.flightClass,
            onView: () => handleViewTicket(ticket),
            onDownload: () => handleDownload(ticketRef.current),
          }))}
          ticketRef={ticketRef}
        />
      )}
      <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
        <DialogOverlay />
        <DialogContent>
          {selectedTicket && selectedTicket.status === 'Active' ? (
            <div ref={ticketRef}>
              <ModernFlightTicket
                passengerName={`${selectedTicket.owner.first_name} ${selectedTicket.owner.last_name}`}
                flightNumber={departureFlightData.flightNumber}
                flightDate={new Date(departureFlightData.departureTime.seconds * 1000).toLocaleDateString('vi-VN', { day: '2-digit', month: '2-digit', year: 'numeric' })}
                departureTime={new Date(departureFlightData.departureTime.seconds * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                from={departureFlightData.departureCity}
                to={departureFlightData.arrivalCity}
                flightClass={selectedTicket.flightClass}
                boardingTime={new Date(departureFlightData.departureTime.seconds * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                gate={6}
                seat={selectedTicket.seatCode || 'N/A'}
              />
              <div className="mt-4 flex justify-end gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => handleDownload(ticketRef.current)}
                  className="gap-2"
                >
                  <Download className="w-4 h-4" />
                  Tải về
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => handleCancelTicket(selectedTicket.originalId)}
                  className="gap-2"
                >
                  Hủy vé
                </Button>
              </div>
            </div>
          ) : (
            <div>
              <p>Vé không còn khả dụng hoặc không ở trạng thái Active.</p>
            </div>
          )}
        </DialogContent>
      </Dialog>

      <div className="flex justify-center mt-8">
        <Button variant="outline" onClick={() => router.push('/')}>Quay trở lại trang chủ</Button>
      </div>
    </div>
  );
}
