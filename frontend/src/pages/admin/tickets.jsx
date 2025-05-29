'use client'

import { useState, useEffect } from 'react'
import { Button } from "@/components/ui/button"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from "@/components/ui/dialog"
import { Badge } from "@/components/ui/badge"
import { useRouter } from 'next/router'
import { toast } from '@/hooks/use-toast'

export default function FlightManagement() {
  const router = useRouter()

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) {
      router.push('/admin')
    }
    else getAllFlights()

  }, [router])

  const [flights, setFlights] = useState([])
  const [selectedFlight, setSelectedFlight] = useState(null)
  const [tickets, setTickets] = useState([])
  const [isDialogOpen, setIsDialogOpen] = useState(false)

  const handleFlightClick = (flight) => {
    setSelectedFlight(flight)
    getTickets(flight)
  }

  const getAllFlights = async () => {
    const getAllFlightsApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/flight/all`

    try {
        const response = await fetch(getAllFlightsApi, {
            method: "GET",
            headers: {
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            },
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }

        const res = await response.json()
        setFlights(res.data.flights.map(a => {return {
            "id": a.flight_id,
            "flightNumber": a.flight_number,
            "arrival": `${a.arrival_city} ${new Date(a.arrival_time.seconds * 1000).toISOString().replace("T", " ").slice(0, -5)}`,
            "departure": `${a.departure_city} ${new Date(a.departure_time.seconds * 1000).toISOString().replace("T", " ").slice(0, -5)}`,
            "ticketList": a.ticket_list,
          }}))
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }

  const getTickets = async (flight) => {
    const getTicketsApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/ticket/list?flightId=${flight.id}`

    if(flight.ticketList.length === 0){
      setTickets([])
      setIsDialogOpen(true)
    }
    else {
      try {
        const response = await fetch(getTicketsApi, {
            method: "GET",
            headers: {
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            },
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }

        const res = await response.json()
        setTickets(res.map(a => {return {
          "status": a.status,
          "seatCode": a.seatCode,
          "updatedAt": a.updatedAt,
          "ownerData": a.owner,
          "bookingId": a.bookingId,
          "flightClass": a.flightClass,
          "ticketId": a.ticketId,
          "price": a.price,
          "flightId": a.flightId,
          "createdAt": a.createdAt
        }}))
        setIsDialogOpen(true)
      } catch (error) {
        toast({
          title: "Tải thông tin vé không thành công",
          description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
          variant: "destructive"
        })
      }
    }
  }

  const handleCancelTicket = async (ticketId) => {
    const cancelTicketApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/ticket/cancel?id=${ticketId}`

    try {
        const response = await fetch(cancelTicketApi, {
            method: "PUT",
            headers: {
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            },
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }

        const res = await response.json()
        toast({
          title: "Thành công",
          description: "Vé đã được hủy thành công",
        })
        setIsDialogOpen(false)
    } catch (error) {
      toast({
        title: "Hủy vé không thành công",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }
  console.log(flights)
  return (
    <div className="container mx-auto pt-10 pl-64">
      <h1 className="text-2xl font-semibold mb-10">Quản Lý Đặt Vé</h1>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Số hiệu</TableHead>
            <TableHead>Cất cánh</TableHead>
            <TableHead>Hạ cánh</TableHead>
            <TableHead>Số vé</TableHead>
            <TableHead>Hành động</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {flights && flights.length > 0 ? (
            flights.map((flight) => (
              <TableRow key={flight.id}>
                <TableCell>{flight.flightNumber}</TableCell>
                <TableCell>{flight.departure}</TableCell>
                <TableCell>{flight.arrival}</TableCell>
                <TableCell>{flight.ticketList?.length || 0}</TableCell>
                <TableCell>
                  <Button
                    className="bg-green-600 hover:bg-green-500"
                    onClick={() => handleFlightClick(flight)}
                  >
                    Xem vé
                  </Button>
                </TableCell>
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={5} className="text-center">
                Đang tải dữ liệu...
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>

      <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
        <DialogContent className="max-w-4xl max-h-[80vh] overflow-hidden flex flex-col">
          <DialogHeader>
            <DialogTitle>Danh sách vé - Chuyến bay {selectedFlight?.flightNumber}</DialogTitle>
            <DialogDescription>
              Cất cánh: {selectedFlight && selectedFlight.departure} -
              Hạ cánh: {selectedFlight && selectedFlight.arrival}
            </DialogDescription>
          </DialogHeader>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Mã vé</TableHead>
                <TableHead>Trạng thái</TableHead>
                <TableHead>Ghế</TableHead>
                <TableHead>Hạng</TableHead>
                <TableHead>Giá</TableHead>
                <TableHead>Thông tin hành khách</TableHead>
                <TableHead>Thao tác</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {tickets.map((ticket) => (
                <TableRow key={ticket.ticketId}>
                  <TableCell>{ticket.ticketId}</TableCell>
                  <TableCell>
                    <Badge variant={ticket.status === 'Active' ? 'success' : 'secondary'}>
                      {ticket.status}
                    </Badge>
                  </TableCell>
                  <TableCell>{ticket.seatCode}</TableCell>
                  <TableCell>{ticket.flightClass}</TableCell>
                  <TableCell>{ticket.price.toLocaleString('vi-VN')} VND</TableCell>
                  <TableCell>
                    <p>{`${ticket.ownerData.firstName} ${ticket.ownerData.lastName}`}</p>
                    <p>{ticket.ownerData.phoneNumber}</p>
                  </TableCell>
                  <TableCell>
                    <Button
                      className="bg-red-500 hover:bg-red-600"
                      onClick={() => handleCancelTicket(ticket.ticketId)}
                    >
                      Xóa
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </DialogContent>
      </Dialog>
    </div>
  )
}

