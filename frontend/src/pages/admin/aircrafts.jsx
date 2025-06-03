"use client"

import { useEffect, useState } from "react"
import { Search, Plus } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { useRouter } from "next/router"
import { toast } from "@/hooks/use-toast"

// Dữ liệu cứng về các tàu bay
const mockAircrafts = [
  {
    id: "AC001",
    name: "Boeing 737-800",
    manufacturer: "Boeing",
    yearOfManufacture: 2018,
    totalSeats: 189,
    seatConfig: {
      economy: 162,
      business: 27
    },
    status: "Active"
  },
  {
    id: "AC002",
    name: "Airbus A320",
    manufacturer: "Airbus",
    yearOfManufacture: 2019,
    totalSeats: 180,
    seatConfig: {
      economy: 150,
      business: 30
    },
    status: "Active"
  },
  {
    id: "AC003",
    name: "Boeing 787-9",
    manufacturer: "Boeing",
    yearOfManufacture: 2020,
    totalSeats: 290,
    seatConfig: {
      economy: 240,
      business: 50
    },
    status: "Maintenance"
  }
]

export default function AircraftManagement() {
  const router = useRouter()
  const [aircrafts, setAircrafts] = useState(mockAircrafts)
  const [searchQuery, setSearchQuery] = useState("")
  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false)
  const [newAircraft, setNewAircraft] = useState({
    id: "",
    name: "",
    manufacturer: "",
    yearOfManufacture: "",
    totalSeats: "",
    seatConfig: {
      economy: "",
      business: ""
    },
    status: "Active"
  })

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) {
      router.push('/admin')
    }
  }, [router])

  const handleAddAircraft = () => {
    // Kiểm tra dữ liệu
    if (!newAircraft.id || !newAircraft.name || !newAircraft.manufacturer || 
        !newAircraft.yearOfManufacture || !newAircraft.totalSeats) {
      toast({
        title: "Lỗi",
        description: "Vui lòng điền đầy đủ thông tin",
        variant: "destructive"
      })
      return
    }

    // Thêm tàu bay mới vào danh sách
    setAircrafts([...aircrafts, newAircraft])
    setIsAddDialogOpen(false)
    setNewAircraft({
      id: "",
      name: "",
      manufacturer: "",
      yearOfManufacture: "",
      totalSeats: "",
      seatConfig: {
        economy: "",
        business: ""
      },
      status: "Active"
    })
    toast({
      title: "Thành công",
      description: "Đã thêm tàu bay mới",
    })
  }

  const handleDeleteAircraft = (id) => {
    setAircrafts(aircrafts.filter(aircraft => aircraft.id !== id))
    toast({
      title: "Thành công",
      description: "Đã xóa tàu bay",
    })
  }

  const filteredAircrafts = aircrafts.filter(
    aircraft =>
      aircraft.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
      aircraft.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      aircraft.manufacturer.toLowerCase().includes(searchQuery.toLowerCase())
  )

  return (
    <div className="pt-10 pl-64 mx-auto">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-semibold">Quản Lý Tàu Bay</h1>
        <Dialog open={isAddDialogOpen} onOpenChange={setIsAddDialogOpen}>
          <DialogTrigger asChild>
            <Button className="bg-orange hover:bg-orange/90">
              <Plus className="mr-2 h-4 w-4" />
              Thêm tàu bay
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Thêm tàu bay mới</DialogTitle>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-4 items-center gap-4">
                <label htmlFor="id" className="text-right">Mã tàu bay</label>
                <Input
                  id="id"
                  value={newAircraft.id}
                  onChange={(e) => setNewAircraft({...newAircraft, id: e.target.value})}
                  className="col-span-3"
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <label htmlFor="name" className="text-right">Tên tàu bay</label>
                <Input
                  id="name"
                  value={newAircraft.name}
                  onChange={(e) => setNewAircraft({...newAircraft, name: e.target.value})}
                  className="col-span-3"
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <label htmlFor="manufacturer" className="text-right">Hãng sản xuất</label>
                <Input
                  id="manufacturer"
                  value={newAircraft.manufacturer}
                  onChange={(e) => setNewAircraft({...newAircraft, manufacturer: e.target.value})}
                  className="col-span-3"
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <label htmlFor="year" className="text-right">Năm sản xuất</label>
                <Input
                  id="year"
                  type="number"
                  value={newAircraft.yearOfManufacture}
                  onChange={(e) => setNewAircraft({...newAircraft, yearOfManufacture: e.target.value})}
                  className="col-span-3"
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <label htmlFor="totalSeats" className="text-right">Tổng số ghế</label>
                <Input
                  id="totalSeats"
                  type="number"
                  value={newAircraft.totalSeats}
                  onChange={(e) => setNewAircraft({...newAircraft, totalSeats: e.target.value})}
                  className="col-span-3"
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <label htmlFor="economySeats" className="text-right">Số ghế phổ thông</label>
                <Input
                  id="economySeats"
                  type="number"
                  value={newAircraft.seatConfig.economy}
                  onChange={(e) => setNewAircraft({
                    ...newAircraft, 
                    seatConfig: {...newAircraft.seatConfig, economy: e.target.value}
                  })}
                  className="col-span-3"
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <label htmlFor="businessSeats" className="text-right">Số ghế thương gia</label>
                <Input
                  id="businessSeats"
                  type="number"
                  value={newAircraft.seatConfig.business}
                  onChange={(e) => setNewAircraft({
                    ...newAircraft, 
                    seatConfig: {...newAircraft.seatConfig, business: e.target.value}
                  })}
                  className="col-span-3"
                />
              </div>
            </div>
            <div className="flex justify-end">
              <Button onClick={handleAddAircraft}>Thêm</Button>
            </div>
          </DialogContent>
        </Dialog>
      </div>

      <div className="relative mb-6">
        <Input
          type="text"
          placeholder="Tìm kiếm tàu bay theo mã, tên hoặc hãng sản xuất"
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
              <TableHead className="text-center">MÃ TÀU BAY</TableHead>
              <TableHead className="text-center">TÊN TÀU BAY</TableHead>
              <TableHead className="text-center">HÃNG SẢN XUẤT</TableHead>
              <TableHead className="text-center">NĂM SẢN XUẤT</TableHead>
              <TableHead className="text-center">TỔNG SỐ GHẾ</TableHead>
              <TableHead className="text-center">GHẾ PHỔ THÔNG</TableHead>
              <TableHead className="text-center">GHẾ THƯƠNG GIA</TableHead>
              <TableHead className="text-center">TRẠNG THÁI</TableHead>
              <TableHead className="text-center">THAO TÁC</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredAircrafts.map((aircraft) => (
              <TableRow key={aircraft.id}>
                <TableCell className="text-center">{aircraft.id}</TableCell>
                <TableCell className="text-center">{aircraft.name}</TableCell>
                <TableCell className="text-center">{aircraft.manufacturer}</TableCell>
                <TableCell className="text-center">{aircraft.yearOfManufacture}</TableCell>
                <TableCell className="text-center">{aircraft.totalSeats}</TableCell>
                <TableCell className="text-center">{aircraft.seatConfig.economy}</TableCell>
                <TableCell className="text-center">{aircraft.seatConfig.business}</TableCell>
                <TableCell className="text-center">
                  <span className={`px-2 py-1 rounded-full text-xs ${
                    aircraft.status === 'Active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'
                  }`}>
                    {aircraft.status === 'Active' ? 'Hoạt động' : 'Bảo trì'}
                  </span>
                </TableCell>
                <TableCell className="text-center">
                  <Button
                    variant="destructive"
                    size="sm"
                    onClick={() => handleDeleteAircraft(aircraft.id)}
                    className="bg-red-500 hover:bg-red-600 text-white text-xs px-3 py-1 h-7"
                  >
                    Xóa
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  )
} 