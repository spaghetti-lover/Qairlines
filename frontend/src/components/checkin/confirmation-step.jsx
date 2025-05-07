import { Check, Printer, Download, Mail, ArrowLeft, Home } from 'lucide-react';
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import { useState, useRef } from "react";
import ModernFlightTicket from "./flight-ticket";
import jsPDF from "jspdf";
import html2canvas from "html2canvas";

export function ConfirmationStep({
  bookingReference,
  departurePassengers,
  returnPassengers,
  departureFlight,
  returnFlight,
  onBack,
  onHome,
}) {
  const [showBoardingPass, setShowBoardingPass] = useState(false);
  const [currentTicket, setCurrentTicket] = useState(null); // Lưu vé hiện tại để hiển thị trong Dialog
  const ticketRef = useRef();


  const calculateBoardingTime = (departureTime) => {
    if (!departureTime) return "N/A";
    
    try {
      // Giả sử departureTime dạng "07:45 AM"
      const parts = departureTime.split(" ");
      if (parts.length !== 2) {
        console.error("Invalid departure time format");
        return "N/A";
      }
  
      const [timeStr, meridiem] = parts;
      const [hourStr, minuteStr] = timeStr.split(":");
      if (!hourStr || !minuteStr) {
        console.error("Invalid time format");
        return "N/A";
      }
  
      let hour = parseInt(hourStr, 10);
      const minute = parseInt(minuteStr, 10);
  
      // Chuyển đổi sang 24h
      const meridiemUpper = meridiem.toUpperCase();
      if (meridiemUpper === "PM" && hour !== 12) {
        hour += 12;
      } else if (meridiemUpper === "AM" && hour === 12) {
        hour = 0;
      }
  
      // Tạo đối tượng Date hôm nay với giờ và phút tương ứng
      const now = new Date();
      now.setHours(hour);
      now.setMinutes(minute);
      now.setSeconds(0);
      now.setMilliseconds(0);
  
      // Lấy thời gian dưới dạng giây
      const departureTimeInSeconds = Math.floor(now.getTime() / 1000);
      const boardingTimeInSeconds = departureTimeInSeconds - 30 * 60; // trừ 30 phút
      const boardingDate = new Date(boardingTimeInSeconds * 1000);
  
      if (isNaN(boardingDate.getTime())) {
        console.error("Invalid boarding time:", boardingTimeInSeconds);
        return "N/A";
      }
  
      return boardingDate.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
    } catch (error) {
      console.error("Error calculating boarding time:", error);
      return "N/A";
    }
  };
  
  
  // Tải vé dưới dạng PDF
  const handleDownload = async () => {
    try {
      const ticketElement = ticketRef.current;
      if (!ticketElement) {
        console.error("Ticket element not found!");
        return;
      }

      const canvas = await html2canvas(ticketElement);
      const imgData = canvas.toDataURL("image/png");

      const pdf = new jsPDF();
      const pdfWidth = pdf.internal.pageSize.getWidth();
      const pdfHeight = (canvas.height * pdfWidth) / canvas.width;

      pdf.addImage(imgData, "PNG", 0, 0, pdfWidth, pdfHeight);
      pdf.save("boarding-pass.pdf");
    } catch (error) {
      console.error("Error during download:", error);
    }
  };

  // In vé
  const handlePrint = () => {
    try {
      const ticketElement = ticketRef.current;
      if (!ticketElement) {
        console.error("Ticket element not found!");
        return;
      }

      const printWindow = window.open("", "_blank");
      printWindow.document.write("<html><head><title>Boarding Pass</title></head><body>");
      printWindow.document.write(ticketElement.outerHTML);
      printWindow.document.write("</body></html>");
      printWindow.document.close();
      printWindow.print();
    } catch (error) {
      console.error("Error during print:", error);
    }
  };

  // Hiển thị danh sách vé cho hành khách
  const renderTickets = (passengers, flight, title) => (
    <div className="mb-6">
      <h2 className="text-xl font-semibold text-orange mb-4">{title}</h2>
      {passengers.map((passenger) => (
        <Card key={passenger.id} className="mb-4">
          <CardContent className="p-6">
            <div className="flex justify-between items-center">
              <div>
                <p className="text-sm text-zinc-500">Hành khách</p>
                <p className="text-lg font-semibold">{passenger.name}</p>
              </div>
              <div>
                <p className="text-sm text-zinc-500">Ghế</p>
                <p className="text-lg font-semibold">{passenger.seat || "Chưa chọn"}</p>
              </div>
            </div>
            <div className="mt-4 flex gap-2">
              <Button
                variant="outline"
                size="sm"
                onClick={() => {
                  setCurrentTicket({ passenger, flight });
                  setShowBoardingPass(true);
                }}
                className="gap-2"
              >
                <Printer className="w-4 h-4" />
                Xem vé
              </Button>
              <Button
                variant="outline"
                size="sm"
                onClick={() => {
                  setCurrentTicket({ passenger, flight });
                  handleDownload();
                }}
                className="gap-2"
              >
                <Download className="w-4 h-4" />
                Tải về
              </Button>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );

  return (
    <div className="max-w-4xl mx-auto p-4">
      <h1 className="text-2xl text-orange font-semibold mb-6">Xác nhận làm thủ tục lên máy bay</h1>

      <Card className="mb-6">
        <CardContent className="p-6">
          <div className="flex items-start gap-4">
            <div className="w-10 h-10 rounded-full bg-green-100 flex items-center justify-center flex-shrink-0">
              <Check className="w-6 h-6 text-green-600" />
            </div>
            <div>
              <h2 className="text-lg font-semibold mb-2">Quý khách đã làm thủ tục lên máy bay thành công</h2>
              <p className="text-zinc-600">
                Quý khách đã làm thủ tục thành công. Thẻ lên máy bay của Quý khách nên được lưu trên máy hoặc in ra. Tại
                sân bay, Quý khách vui lòng kiểm tra lại thông tin về cửa ra máy bay.
              </p>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Chuyến đi */}
      {renderTickets(departurePassengers, departureFlight, "Chuyến đi")}

      {/* Chuyến về */}
      {renderTickets(returnPassengers, returnFlight, "Chuyến về")}

      <div className="flex justify-between mt-6">
        <Button variant="outline" onClick={onBack} className="gap-2">
          <ArrowLeft className="w-4 h-4" />
          Quay lại
        </Button>
        <Button variant="orange" onClick={onHome} className="gap-2">
          <Home className="w-4 h-4" />
          Về trang chính
        </Button>
      </div>

      {/* Dialog hiển thị thẻ lên máy bay */}
      <Dialog open={showBoardingPass} onOpenChange={setShowBoardingPass}>
        <DialogContent className="sm:max-w-md">
          <div ref={ticketRef}>
            {currentTicket && (
              <ModernFlightTicket
                passengerName={currentTicket.passenger.name}
                flightNumber={currentTicket.flight.flightNumber}
                flightDate={currentTicket.flight.date}
                departureTime={currentTicket.flight.departureTime}
                from={currentTicket.flight.from}
                to={currentTicket.flight.to}
                flightClass={
                  currentTicket.passenger.type?.trim().toLowerCase() === "business"
                    ? "Thương gia"
                    : "Phổ thông"
                }
                boardingTime={calculateBoardingTime(currentTicket.flight.departureTime)}
                gate={currentTicket.flight.gate}
                seat={currentTicket.passenger.seat || "Chưa chọn"}
              />
            )}
          </div>
          <div className="mt-4 flex justify-end gap-2">
            <Button variant="outline" size="sm" onClick={handleDownload} className="gap-2">
              <Download className="w-4 h-4" />
              Tải về
            </Button>
            <Button variant="outline" size="sm" onClick={handlePrint} className="gap-2">
              <Printer className="w-4 h-4" />
              In vé
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}

