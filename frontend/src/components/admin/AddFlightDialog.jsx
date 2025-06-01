'use client'

import { useState } from 'react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { toast } from "@/hooks/use-toast"
import { Plus } from 'lucide-react'

export function AddFlightDialog() {
  const [isOpen, setIsOpen] = useState(false)
  const [flightData, setFlightData] = useState({
    // flightId: '',
    flightNumber: '',
    aircraftType: '',
    departureCity: '',
    arrivalCity: '',
    departureTime: '',
    arrivalTime: '',
    basePrice: '',
    status: 'On Time'
  })

  const handleInputChange = (e) => {
    const { name, value } = e.target
    setFlightData(prev => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const createFlightApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/flight`
    const formattedFlightData = {
      ...flightData,
      departureTime: new Date(flightData.departureTime).toISOString(),
      arrivalTime: new Date(flightData.arrivalTime).toISOString(),
      basePrice: parseInt(flightData.basePrice, 10),
    };
    try {
        const response = await fetch(createFlightApi, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "admin": "true",
              "authorization": "Bearer " + localStorage.getItem("token")
            },
            body: JSON.stringify(formattedFlightData)
        })
        if (!response.ok) {
            throw new Error("failed")
        }
        toast({
          title: "Đã thêm chuyến bay",
          description: `Chuyến bay ${flightData.flightNumber} đã được thêm thành công.`,
        })
        window.location.reload();
    } catch (error) {
        alert("Đã xảy ra lỗi, vui lòng thử lại")
        console.log(error)
    }
    setIsOpen(false)
    // Reset form after submission
    setFlightData({
      // flightId: '',
      flightNumber: '',
      aircraftType: '',
      departureCity: '',
      arrivalCity: '',
      departureAirport: '',
      arrivalAirport: '',
      departureTime: '',
      arrivalTime: '',
      basePrice: '',
      status: 'On Time'
    })
  }

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button className="bg-yellow-400 hover:bg-yellow-500 text-black font-medium">
          <Plus className="mr-2 h-4 w-4" />
          CHUYẾN BAY MỚI
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Thêm chuyến bay mới</DialogTitle>
          <DialogDescription>
            Vui lòng nhập đầy đủ các thông tin bên dưới
          </DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <div className="grid gap-4">
            {/* <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="flightId" className="text-right">
                ID chuyến bay
              </Label>
              <Input
                id="flightId"
                name="flightId"
                value={flightData.flightId}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div> */}
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="flightNumber" className="text-right">
                Số hiệu chuyến
              </Label>
              <Input
                id="flightNumber"
                name="flightNumber"
                value={flightData.flightNumber}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="aircraftType" className="text-right">
                Loại tàu bay
              </Label>
              <Input
                id="aircraftType"
                name="aircraftType"
                value={flightData.aircraftType}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="departureCity" className="text-right">
                Điểm cất cánh
              </Label>
              <Input
                id="departureCity"
                name="departureCity"
                value={flightData.departureCity}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="arrivalCity" className="text-right">
              Điểm hạ cánh
              </Label>
              <Input
                id="arrivalCity"
                name="arrivalCity"
                value={flightData.arrivalCity}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="departureAirport" className="text-right">
                Sân bay cất cánh
              </Label>
              <Input
                id="departureAirport"
                name="departureAirport"
                value={flightData.departureAirport}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="arrivalAirport" className="text-right">
                Sân bay hạ cánh
              </Label>
              <Input
                id="arrivalAirport"
                name="arrivalAirport"
                value={flightData.arrivalAirport}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="departureTime" className="text-right">
                Thời gian cất cánh
              </Label>
              <Input
                id="departureTime"
                name="departureTime"
                type="datetime-local"
                value={flightData.departureTime}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="arrivalTime" className="text-right">
              Thời gian hạ cánh
              </Label>
              <Input
                id="arrivalTime"
                name="arrivalTime"
                type="datetime-local"
                value={flightData.arrivalTime}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="basePrice" className="text-right">
                Giá cơ sở
              </Label>
              <Input
                id="basePrice"
                name="basePrice"
                type="number"
                step="0.01"
                value={flightData.basePrice}
                onChange={handleInputChange}
                className="col-span-3"
                required
              />
            </div>
          </div>
          <DialogFooter>
            <Button type="submit" className="bg-blue-500 hover:bg-blue-600">Lưu</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}

