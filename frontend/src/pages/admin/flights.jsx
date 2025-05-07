"use client"

import { useEffect, useState } from "react"
import { Search } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import { toast } from "@/hooks/use-toast"

import { useRouter } from "next/router"

import { AddFlightDialog } from "@/components/admin/AddFlightDialog"
import { EditFlightDialog } from "@/components/admin/EditFlightDialog"

export default function ScheduledFlights() {
  const router = useRouter()

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) {
      router.push('/admin')
    }
    else getAllFlights()

  }, [router])

  const [flights, setFlights] = useState([])
  const [searchQuery, setSearchQuery] = useState("")
  const [editingFlight, setEditingFlight] = useState(null)

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
        setFlights(res.data.map(a => {return {
            "flightId": a.flightId,
            "id": a.flightNumber,
            "aircraft": a.aircraftType, 
            "src": a.arrivalCity, 
            "dest": a.departureCity,
            "adt": new Date(a.arrivalTime.seconds*1000).toISOString().replace("T", " ").slice(0, -5), 
            "ddt": new Date(a.departureTime.seconds*1000).toISOString().replace("T", " ").slice(0, -5),
            "cec": a.basePrice*1.5, 
            "cbc": a.basePrice*2, 
            "noe": 198, 
            "nob": 66, 
            "status": a.status
          }}))
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }

  const handleRemove = async (id) => {
    setFlights(flights.filter(flight => flight.id !== id))
    const deleteFlightApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/flight/delete/?`

    try {
        const response = await fetch(deleteFlightApi + 
        new URLSearchParams({
          "id": id,
        }).toString(), {
            method: "DELETE",
            headers: {
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            }, 
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }
    } catch (error) {
      toast({
        title: "Hủy chuyến không thành công",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }

    toast({
      title: "Thông báo",
      description: "Chuyến bay đã được xóa thành công.",
    })
  }

  const handleEditComplete = (updatedFlight) => {
    setFlights(flights.map(f => f.id === updatedFlight.id ? updatedFlight : f))
    setEditingFlight(null)
    toast({
      title: "Thông báo",
      description: "Thông tin chuyến bay đã được cập nhật thành công.",
    })
  }

  const getStatusBadge = (flight) => {
    switch(flight.status) {
      case 'OnTime':
        return <Badge className="bg-green-400 hover:bg-green-400 text-black">Đang Bay</Badge>
      case 'Landed':
        return <Badge className="bg-yellow-400 hover:bg-yellow-400 text-black">Đã Hạ Cánh</Badge>
      default:
        return (
          <div className="flex gap-2">
            <Button 
              size="sm"
              className="bg-cyan-500 hover:bg-cyan-600 text-white text-xs px-3 py-1 h-7"
              onClick={() => setEditingFlight(flight)}
            >
              Sửa
            </Button>
            <Button 
              size="sm"
              variant="destructive" 
              onClick={() => handleRemove(flight.id)}
              className="bg-red-500 hover:bg-red-600 text-white text-xs px-3 py-1 h-7"
            >
              Hủy
            </Button>
          </div>
        )
    }
  }

  return (
    <div className="pt-10 pl-64 mx-auto">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-semibold">Quản Lý Chuyến Bay</h1>
        <AddFlightDialog />

      </div>

      <div className="relative mb-6">
        <Input
          type="text"
          placeholder="Tìm kiếm chuyến bay sử dụng tên tàu bay hoặc số hiệu chuyến"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="pl-4 pr-10 h-10 border rounded"
        />
        <Button 
          size="sm" 
          className="absolute right-0 top-0 h-10 bg-blue-500 hover:bg-blue-600 rounded-l-none"
        >
          <Search className="h-4 w-4" />
        </Button>
      </div>

      <div className="border rounded-sm">
        <Table>
          <TableHeader>
            <TableRow className="bg-gray-100">
              <TableHead className="w-[80px] text-center">SỐ HIỆU</TableHead>
              <TableHead className="text-center">LOẠI MÁY BAY</TableHead>
              <TableHead className="text-center">CẤT CÁNH</TableHead>
              <TableHead className="text-center">HẠ CÁNH</TableHead>
              <TableHead className="text-center">PHỔ THÔNG</TableHead>
              <TableHead className="text-center">THƯƠNG GIA</TableHead>
              <TableHead className="text-center">TÙY CHỈNH</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {flights.filter(
              flight => flight.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
              flight.aircraft.toLowerCase().includes(searchQuery.toLowerCase())
            ).map((flight) => (
              <TableRow key={flight.flightId} className={flight.id % 2 === 0 ? "bg-white" : "bg-gray-50"}>
                <TableCell className="text-center">{flight.id}</TableCell>
                <TableCell className="text-center">{flight.aircraft}</TableCell>
                <TableCell className="text-center">{`${flight.src} ${flight.ddt}`}</TableCell>
                <TableCell className="text-center">{`${flight.dest} ${flight.adt}`}</TableCell>
                <TableCell className="text-center">{`${flight.cec} VND x ${flight.noe}`}</TableCell>
                <TableCell className="text-center">{`${flight.cbc} VND x ${flight.nob}`}</TableCell>
                <TableCell>
                  <div className="flex gap-2">
                    {getStatusBadge(flight)}
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        
      </div>
      {editingFlight && (
        <EditFlightDialog
          flight={editingFlight}
          onClose={() => setEditingFlight(null)}
          onSave={handleEditComplete}
        />
      )}
    </div>
  )
}