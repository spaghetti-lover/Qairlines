import { useEffect, useState, useRef } from 'react';
import axios from 'axios';
import html2canvas from 'html2canvas';
import jsPDF from 'jspdf';
import { useToast } from '@/hooks/use-toast';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export function useFlightBooking(bookingID) {
  const [bookingData, setBookingData] = useState(null);
  const [departureTicketData, setDepartureTicketData] = useState([]);
  const [returnTicketData, setReturnTicketData] = useState([]);
  const [departureFlightData, setDepartureFlightData] = useState(null);
  const [returnFlightData, setReturnFlightData] = useState(null);
  const [selectedTicket, setSelectedTicket] = useState(null);
  const [dialogOpen, setDialogOpen] = useState(false);
  const [error, setError] = useState(null);
  const ticketRef = useRef(null);

  const { toast } = useToast();

  useEffect(() => {
    if (!bookingID) return;
    const token = localStorage.getItem('token');

    if (!token) {
      setError('Không tìm thấy token. Vui lòng đăng nhập lại.');
      return;
    }

    // Fetch booking data
    axios
      .get(`${API_BASE_URL}/api/booking/?id=${bookingID}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((response) => {
        const booking = response.data.data;
        setBookingData(booking);

        // Fetch departure tickets data
        if (booking.departureIdTickets && booking.departureIdTickets.length > 0) {
          Promise.all(
            booking.departureIdTickets.map((ticketID) =>
              axios.get(`${API_BASE_URL}/api/ticket?id=${ticketID}`, {
                headers: {
                  Authorization: `Bearer ${token}`,
                },
              })
            )
          )
            .then((responses) => {
              // Gán originalId tương ứng với từng vé
              const departureTickets = responses.map((res, index) => ({
                ...res.data.data,
                originalId: booking.departureIdTickets[index],
              }));
              setDepartureTicketData(departureTickets);
            })
            .catch((err) => {
              setError('Không thể tải dữ liệu vé đi. Vui lòng thử lại sau.');
              console.error(err);
            });
        }

        // Fetch return tickets data
        if (booking.returnIdTickets && booking.returnIdTickets.length > 0) {
          Promise.all(
            booking.returnIdTickets.map((ticketID) =>
              axios.get(`${API_BASE_URL}/api/ticket?id=${ticketID}`, {
                headers: {
                  Authorization: `Bearer ${token}`,
                },
              })
            )
          )
            .then((responses) => {
              const returnTickets = responses.map((res, index) => ({
                ...res.data.data,
                originalId: booking.returnIdTickets[index],
              }));
              setReturnTicketData(returnTickets);
            })
            .catch((err) => {
              setError('Không thể tải dữ liệu vé về. Vui lòng thử lại sau.');
              console.error(err);
            });
        }

        // Fetch departure flight data
        if (booking.departureFlightId) {
          axios
            .get(`${API_BASE_URL}/api/flight?id=${booking.departureFlightId}`, {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            })
            .then((response) => {
              setDepartureFlightData(response.data.data);
            })
            .catch((err) => {
              setError('Không thể tải dữ liệu chuyến bay đi. Vui lòng thử lại sau.');
              console.error(err);
            });
        }

        // Fetch return flight data
        if (booking.returnFlightId) {
          axios
            .get(`${API_BASE_URL}/api/flight?id=${booking.returnFlightId}`, {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            })
            .then((response) => {
              setReturnFlightData(response.data.data);
            })
            .catch((err) => {
              setError('Không thể tải dữ liệu chuyến bay về. Vui lòng thử lại sau.');
              console.error(err);
            });
        }
      })
      .catch((err) => {
        setError('Không thể tải dữ liệu đặt chỗ. Vui lòng thử lại sau.');
        console.error(err);
      });
  }, [bookingID]);

  const handleViewTicket = (ticket) => {
    const translatedClass =
      ticket.flightClass === 'economy'
        ? 'Phổ thông'
        : ticket.flightClass === 'business'
        ? 'Thương gia'
        : 'N/A';
    setSelectedTicket({
      ...ticket,
      flightClass: translatedClass,
      gate: 6
    });
    setDialogOpen(true);
  };

  const handleDownload = async (ticketElement) => {
    try {
      if (!ticketElement) {
        console.error('Ticket element not found!');
        return;
      }

      const canvas = await html2canvas(ticketElement);
      const imgData = canvas.toDataURL('image/png');

      const pdf = new jsPDF();
      const pdfWidth = pdf.internal.pageSize.getWidth();
      const pdfHeight = (canvas.height * pdfWidth) / canvas.width;

      pdf.addImage(imgData, 'PNG', 0, 0, pdfWidth, pdfHeight);
      pdf.save('boarding-pass.pdf');
    } catch (error) {
      console.error('Error during download:', error);
    }
  };

  const handleCancelTicket = async (ticketID) => {
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        console.error('Không tìm thấy token. Vui lòng đăng nhập lại.');
        return;
      }

      console.log('Hủy vé với ID:', ticketID);

      // Gọi API PUT để hủy vé
      await axios.put(`${API_BASE_URL}/api/ticket/cancel?id=${ticketID}`, {}, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      // Thông báo thành công
      toast({
        title: "Hủy vé thành công",
        description: "Vé của bạn đã được hủy.",
        variant: "default"
      });

      setDialogOpen(false);

      // Cập nhật state: Đổi status vé bị hủy thành 'cancelled'
      setDepartureTicketData(prev =>
        prev.map((t) => (t.originalId === ticketID ? { ...t, status: 'cancelled' } : t))
      );

      setReturnTicketData(prev =>
        prev.map((t) => (t.originalId === ticketID ? { ...t, status: 'cancelled' } : t))
      );

      if (selectedTicket && selectedTicket.originalId === ticketID) {
        setSelectedTicket({ ...selectedTicket, status: 'cancelled' });
      }

    } catch (error) {
      console.error('Lỗi khi hủy vé:', error);
      if (error.response) {
        console.error('Backend trả về lỗi:', error.response.data);
      }

      // Thông báo lỗi
      toast({
        title: "Không thể hủy vé",
        description: "Đã xảy ra lỗi, vui lòng thử lại sau.",
        variant: "destructive"
      });
    }
  };

  return {
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
  };
}
