"use client"

import { useState } from "react"
import { Search, Plus, Edit, Trash2, MapPin, AlertCircle } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger, DialogFooter } from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { toast } from "@/hooks/use-toast"
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"

const hardcodedAirports = [
  {
    id: 1,
    code: "HAN",
    iataCode: "HAN",
    icaoCode: "VVNB",
    name: "Sân bay Quốc tế Nội Bài",
    city: "Hà Nội",
    country: "Việt Nam",
    terminals: 2,
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: 2,
  },
  {
    id: 2,
    code: "SGN",
    iataCode: "SGN",
    icaoCode: "VVTS",
    name: "Sân bay Quốc tế Tân Sơn Nhất",
    city: "Hồ Chí Minh",
    country: "Việt Nam",
    terminals: 2,
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: 2,
  },
  {
    id: 3,
    code: "DAD",
    iataCode: "DAD",
    icaoCode: "VVDN",
    name: "Sân bay Quốc tế Đà Nẵng",
    city: "Đà Nẵng",
    country: "Việt Nam",
    terminals: 1,
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: 1,
  },
  {
    id: 4,
    code: "CXR",
    iataCode: "CXR",
    icaoCode: "VVCR",
    name: "Sân bay Quốc tế Cam Ranh",
    city: "Nha Trang",
    country: "Việt Nam",
    terminals: 2,
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: 1,
  },
  {
    id: 5,
    code: "PQC",
    iataCode: "PQC",
    icaoCode: "VVPQ",
    name: "Sân bay Quốc tế Phú Quốc",
    city: "Phú Quốc",
    country: "Việt Nam",
    terminals: 1,
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: 1,
  },
  {
    id: 6,
    code: "HPH",
    iataCode: "HPH",
    icaoCode: "VVCI",
    name: "Sân bay Quốc tế Cát Bi",
    city: "Hải Phòng",
    country: "Việt Nam",
    terminals: 1,
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: 1,
  },
  {
    id: 7,
    code: "VCS",
    iataCode: "VCS",
    icaoCode: "VVCS",
    name: "Sân bay Quốc tế Côn Đảo",
    city: "Côn Đảo",
    country: "Việt Nam",
    terminals: 1,
    status: "Bảo trì",
    timezone: "GMT+7",
    runways: 1,
  },
  {
    id: 8,
    code: "VCA",
    iataCode: "VCA",
    icaoCode: "VVCT",
    name: "Sân bay Quốc tế Cần Thơ",
    city: "Cần Thơ",
    country: "Việt Nam",
    terminals: 1,
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: 1,
  },
]

const countries = [
  { value: "Việt Nam", label: "Việt Nam" },
  { value: "Thái Lan", label: "Thái Lan" },
  { value: "Singapore", label: "Singapore" },
  { value: "Malaysia", label: "Malaysia" },
  { value: "Indonesia", label: "Indonesia" },
  { value: "Philippines", label: "Philippines" },
]

const timezones = [
  { value: "GMT+7", label: "GMT+7" },
  { value: "GMT+8", label: "GMT+8" },
  { value: "GMT+9", label: "GMT+9" },
]

const statuses = [
  { value: "Hoạt động", label: "Hoạt động" },
  { value: "Bảo trì", label: "Bảo trì" },
  { value: "Đóng cửa", label: "Đóng cửa" },
]

export default function AirportManagement() {
  const [airports, setAirports] = useState(hardcodedAirports)
  const [searchQuery, setSearchQuery] = useState("")
  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false)
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false)
  const [deleteAirportId, setDeleteAirportId] = useState(null)
  const [editingAirport, setEditingAirport] = useState(null)
  const [errors, setErrors] = useState({})
  const [newAirport, setNewAirport] = useState({
    code: "",
    iataCode: "",
    icaoCode: "",
    name: "",
    city: "",
    country: "Việt Nam",
    terminals: "",
    status: "Hoạt động",
    timezone: "GMT+7",
    runways: "",
  })

  const getStatusBadge = (status) => {
    switch (status) {
      case "Hoạt động":
        return <Badge className="bg-green-100 text-green-800 hover:bg-green-100">Hoạt động</Badge>
      case "Bảo trì":
        return <Badge className="bg-yellow-100 text-yellow-800 hover:bg-yellow-100">Bảo trì</Badge>
      case "Đóng cửa":
        return <Badge className="bg-red-100 text-red-800 hover:bg-red-100">Đóng cửa</Badge>
      default:
        return <Badge variant="secondary">{status}</Badge>
    }
  }

  const safeParseInt = (value) => {
    const parsed = Number.parseInt(value)
    return isNaN(parsed) ? 0 : parsed
  }

  const validateForm = (airport, isEdit = false) => {
    const newErrors = {}

    // Validate mã sân bay
    if (!airport.code.trim()) {
      newErrors.code = "Mã sân bay là bắt buộc"
    } else if (!/^[A-Z]{3}$/.test(airport.code.trim())) {
      newErrors.code = "Mã sân bay phải có 3 ký tự viết hoa"
    } else {
      // Kiểm tra trùng lặp mã sân bay
      const existingAirport = airports.find(
        (a) => a.code.toLowerCase() === airport.code.trim().toLowerCase() && (!isEdit || a.id !== editingAirport?.id),
      )
      if (existingAirport) {
        newErrors.code = "Mã sân bay đã tồn tại"
      }
    }

    // Validate mã IATA
    if (!airport.iataCode.trim()) {
      newErrors.iataCode = "Mã IATA là bắt buộc"
    } else if (!/^[A-Z]{3}$/.test(airport.iataCode.trim())) {
      newErrors.iataCode = "Mã IATA phải có 3 ký tự viết hoa"
    }

    // Validate mã ICAO
    if (!airport.icaoCode.trim()) {
      newErrors.icaoCode = "Mã ICAO là bắt buộc"
    } else if (!/^[A-Z]{4}$/.test(airport.icaoCode.trim())) {
      newErrors.icaoCode = "Mã ICAO phải có 4 ký tự viết hoa"
    }

    // Validate tên sân bay
    if (!airport.name.trim()) {
      newErrors.name = "Tên sân bay là bắt buộc"
    } else if (airport.name.trim().length < 5) {
      newErrors.name = "Tên sân bay phải có ít nhất 5 ký tự"
    }

    // Validate thành phố
    if (!airport.city.trim()) {
      newErrors.city = "Thành phố là bắt buộc"
    }

    // Validate quốc gia
    if (!airport.country) {
      newErrors.country = "Quốc gia là bắt buộc"
    }

    // Validate số nhà ga
    if (airport.terminals && !/^\d+$/.test(airport.terminals.toString().trim())) {
      newErrors.terminals = "Số nhà ga phải là số nguyên dương"
    } else if (airport.terminals && safeParseInt(airport.terminals) <= 0) {
      newErrors.terminals = "Số nhà ga phải lớn hơn 0"
    }

    // Validate số đường băng
    if (airport.runways && !/^\d+$/.test(airport.runways.toString().trim())) {
      newErrors.runways = "Số đường băng phải là số nguyên dương"
    } else if (airport.runways && safeParseInt(airport.runways) <= 0) {
      newErrors.runways = "Số đường băng phải lớn hơn 0"
    }

    return newErrors
  }

  const handleInputChange = (field, value) => {
    // Validation real-time cho các trường số
    if (["terminals", "runways"].includes(field)) {
      // Chỉ cho phép số
      if (value && !/^\d*$/.test(value)) {
        return // Không cho phép nhập ký tự không phải số
      }
    }

    // Auto uppercase cho các mã
    if (["code", "iataCode", "icaoCode"].includes(field)) {
      value = value.toUpperCase()
    }

    setNewAirport({ ...newAirport, [field]: value })

    // Clear error khi user bắt đầu sửa
    if (errors[field]) {
      setErrors({ ...errors, [field]: null })
    }
  }

  const handleAdd = () => {
    const validationErrors = validateForm(newAirport)
    setErrors(validationErrors)

    if (Object.keys(validationErrors).length > 0) {
      toast({
        title: "Lỗi validation",
        description: "Vui lòng kiểm tra lại thông tin đã nhập.",
        variant: "destructive",
      })
      return
    }

    const maxId = airports.length > 0 ? Math.max(...airports.map((a) => a.id || 0)) : 0

    const airport = {
      id: maxId + 1,
      code: newAirport.code.trim().toUpperCase(),
      iataCode: newAirport.iataCode.trim().toUpperCase(),
      icaoCode: newAirport.icaoCode.trim().toUpperCase(),
      name: newAirport.name.trim(),
      city: newAirport.city.trim(),
      country: newAirport.country,
      terminals: safeParseInt(newAirport.terminals),
      status: newAirport.status,
      timezone: newAirport.timezone,
      runways: safeParseInt(newAirport.runways),
    }

    setAirports([...airports, airport])
    setNewAirport({
      code: "",
      iataCode: "",
      icaoCode: "",
      name: "",
      city: "",
      country: "Việt Nam",
      terminals: "",
      status: "Hoạt động",
      timezone: "GMT+7",
      runways: "",
    })
    setErrors({})
    setIsAddDialogOpen(false)
    toast({
      title: "Thành công",
      description: "Thêm sân bay mới thành công.",
    })
  }

  const handleEdit = (airport) => {
    setEditingAirport(airport)
    setNewAirport({
      code: airport.code || "",
      iataCode: airport.iataCode || "",
      icaoCode: airport.icaoCode || "",
      name: airport.name || "",
      city: airport.city || "",
      country: airport.country || "Việt Nam",
      terminals: (airport.terminals || "").toString(),
      status: airport.status || "Hoạt động",
      timezone: airport.timezone || "GMT+7",
      runways: (airport.runways || "").toString(),
    })
    setErrors({})
  }

  const handleUpdate = () => {
    const validationErrors = validateForm(newAirport, true)
    setErrors(validationErrors)

    if (Object.keys(validationErrors).length > 0) {
      toast({
        title: "Lỗi validation",
        description: "Vui lòng kiểm tra lại thông tin đã nhập.",
        variant: "destructive",
      })
      return
    }

    const updatedAirports = airports.map((airport) =>
      airport.id === editingAirport.id
        ? {
            ...airport,
            code: newAirport.code.trim().toUpperCase(),
            iataCode: newAirport.iataCode.trim().toUpperCase(),
            icaoCode: newAirport.icaoCode.trim().toUpperCase(),
            name: newAirport.name.trim(),
            city: newAirport.city.trim(),
            country: newAirport.country,
            terminals: safeParseInt(newAirport.terminals),
            status: newAirport.status,
            timezone: newAirport.timezone,
            runways: safeParseInt(newAirport.runways),
          }
        : airport,
    )

    setAirports(updatedAirports)
    setEditingAirport(null)
    setNewAirport({
      code: "",
      iataCode: "",
      icaoCode: "",
      name: "",
      city: "",
      country: "Việt Nam",
      terminals: "",
      status: "Hoạt động",
      timezone: "GMT+7",
      runways: "",
    })
    setErrors({})
    toast({
      title: "Thành công",
      description: "Cập nhật thông tin sân bay thành công.",
    })
  }

  const confirmDelete = (id) => {
    setDeleteAirportId(id)
    setIsDeleteDialogOpen(true)
  }

  const handleDelete = () => {
    setAirports(airports.filter((airport) => airport.id !== deleteAirportId))
    setIsDeleteDialogOpen(false)
    setDeleteAirportId(null)
    toast({
      title: "Thành công",
      description: "Xóa sân bay thành công.",
    })
  }

  const filteredAirports = airports.filter(
    (airport) =>
      airport.code.toLowerCase().includes(searchQuery.toLowerCase()) ||
      airport.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      airport.city.toLowerCase().includes(searchQuery.toLowerCase()) ||
      airport.country.toLowerCase().includes(searchQuery.toLowerCase()),
  )

  const stats = {
    total: airports.length,
    active: airports.filter((a) => a.status === "Hoạt động").length,
    maintenance: airports.filter((a) => a.status === "Bảo trì").length,
    closed: airports.filter((a) => a.status === "Đóng cửa").length,
  }

  const renderFormField = (
    id,
    label,
    value,
    onChange,
    type = "text",
    required = false,
    placeholder = "",
    options = null,
    maxLength = null,
  ) => (
    <div>
      <Label htmlFor={id} className={required ? "after:content-['*'] after:text-red-500 after:ml-1" : ""}>
        {label}
      </Label>
      {options ? (
        <Select value={value} onValueChange={onChange}>
          <SelectTrigger className={errors[id] ? "border-red-500" : ""}>
            <SelectValue placeholder={placeholder} />
          </SelectTrigger>
          <SelectContent>
            {options.map((option) => (
              <SelectItem key={option.value} value={option.value}>
                {option.label}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      ) : (
        <Input
          id={id}
          type="text"
          inputMode={type === "number" ? "numeric" : "text"}
          pattern={type === "number" ? "[0-9]*" : undefined}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          placeholder={placeholder}
          className={errors[id] ? "border-red-500" : ""}
          maxLength={maxLength}
        />
      )}
      {errors[id] && (
        <div className="flex items-center gap-1 mt-1 text-red-500 text-sm">
          <AlertCircle className="h-3 w-3" />
          {errors[id]}
        </div>
      )}
    </div>
  )

  return (
    <div className="pt-10 pl-64 mx-auto">
      {/* Header */}
      <div className="flex justify-between items-center mb-6">
        <div className="flex items-center gap-3">
          <MapPin className="h-8 w-8 text-blue-600" />
          <h1 className="text-3xl font-bold text-gray-900">Quản Lý Sân Bay</h1>
        </div>
        <Dialog open={isAddDialogOpen} onOpenChange={setIsAddDialogOpen}>
          <DialogTrigger asChild>
            <Button className="bg-blue-600 hover:bg-blue-700 text-white">
              <Plus className="h-4 w-4 mr-2" />
              Thêm Sân Bay Mới
            </Button>
          </DialogTrigger>
          <DialogContent className="max-w-2xl max-h-[90vh] overflow-y-auto">
            <DialogHeader>
              <DialogTitle>Thêm Sân Bay Mới</DialogTitle>
            </DialogHeader>
            <div className="grid grid-cols-2 gap-4">
              {renderFormField(
                "code",
                "Mã Sân Bay",
                newAirport.code,
                (value) => handleInputChange("code", value),
                "text",
                true,
                "HAN",
                null,
                3,
              )}
              {renderFormField(
                "iataCode",
                "Mã IATA",
                newAirport.iataCode,
                (value) => handleInputChange("iataCode", value),
                "text",
                true,
                "HAN",
                null,
                3,
              )}
              {renderFormField(
                "icaoCode",
                "Mã ICAO",
                newAirport.icaoCode,
                (value) => handleInputChange("icaoCode", value),
                "text",
                true,
                "VVNB",
                null,
                4,
              )}
              {renderFormField(
                "name",
                "Tên Sân Bay",
                newAirport.name,
                (value) => handleInputChange("name", value),
                "text",
                true,
                "Sân bay Quốc tế Nội Bài",
              )}
              {renderFormField(
                "city",
                "Thành Phố",
                newAirport.city,
                (value) => handleInputChange("city", value),
                "text",
                true,
                "Hà Nội",
              )}
              {renderFormField(
                "country",
                "Quốc Gia",
                newAirport.country,
                (value) => handleInputChange("country", value),
                "text",
                true,
                "Chọn quốc gia",
                countries,
              )}
              {renderFormField(
                "terminals",
                "Số Nhà Ga",
                newAirport.terminals,
                (value) => handleInputChange("terminals", value),
                "number",
                false,
                "2",
              )}
              {renderFormField(
                "runways",
                "Số Đường Băng",
                newAirport.runways,
                (value) => handleInputChange("runways", value),
                "number",
                false,
                "2",
              )}
              {renderFormField(
                "timezone",
                "Múi Giờ",
                newAirport.timezone,
                (value) => handleInputChange("timezone", value),
                "text",
                false,
                "Chọn múi giờ",
                timezones,
              )}
              {renderFormField(
                "status",
                "Trạng Thái",
                newAirport.status,
                (value) => handleInputChange("status", value),
                "text",
                false,
                "Chọn trạng thái",
                statuses,
              )}
            </div>
            <DialogFooter className="mt-4">
              <Button
                variant="outline"
                onClick={() => {
                  setIsAddDialogOpen(false)
                  setErrors({})
                  setNewAirport({
                    code: "",
                    iataCode: "",
                    icaoCode: "",
                    name: "",
                    city: "",
                    country: "Việt Nam",
                    terminals: "",
                    status: "Hoạt động",
                    timezone: "GMT+7",
                    runways: "",
                  })
                }}
              >
                Hủy
              </Button>
              <Button onClick={handleAdd} className="bg-blue-600 hover:bg-blue-700">
                Thêm Sân Bay
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-gray-600">Tổng Số Sân Bay</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-blue-600">{stats.total}</div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-gray-600">Đang Hoạt Động</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-green-600">{stats.active}</div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-gray-600">Đang Bảo Trì</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-yellow-600">{stats.maintenance}</div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-gray-600">Đóng Cửa</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-red-600">{stats.closed}</div>
          </CardContent>
        </Card>
      </div>

      {/* Search */}
      <div className="relative mb-6">
        <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-gray-400" />
        <Input
          type="text"
          placeholder="Tìm kiếm theo mã, tên, thành phố hoặc quốc gia..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="pl-10 h-12"
        />
      </div>

      {/* Table */}
      <Card className="flex-1">
        <CardContent className="p-0">
          <div className="overflow-x-auto">
            <div className="min-h-[400px]">
              <Table>
                <TableHeader>
                  <TableRow className="bg-gray-50">
                    <TableHead className="text-center font-semibold">STT</TableHead>
                    <TableHead className="text-center font-semibold">Mã</TableHead>
                    <TableHead className="text-center font-semibold">IATA</TableHead>
                    <TableHead className="text-center font-semibold">ICAO</TableHead>
                    <TableHead className="text-center font-semibold">Tên Sân Bay</TableHead>
                    <TableHead className="text-center font-semibold">Thành Phố</TableHead>
                    <TableHead className="text-center font-semibold">Quốc Gia</TableHead>
                    <TableHead className="text-center font-semibold">Nhà Ga</TableHead>
                    <TableHead className="text-center font-semibold">Đường Băng</TableHead>
                    <TableHead className="text-center font-semibold">Trạng Thái</TableHead>
                    <TableHead className="text-center font-semibold">Thao Tác</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {filteredAirports.map((airport, index) => (
                    <TableRow key={airport.id} className={index % 2 === 0 ? "bg-white" : "bg-gray-50"}>
                      <TableCell className="text-center font-medium">{index + 1}</TableCell>
                      <TableCell className="text-center font-semibold text-blue-600">{airport.code}</TableCell>
                      <TableCell className="text-center font-medium">{airport.iataCode}</TableCell>
                      <TableCell className="text-center font-medium">{airport.icaoCode}</TableCell>
                      <TableCell className="font-medium">{airport.name}</TableCell>
                      <TableCell className="text-center font-medium">{airport.city}</TableCell>
                      <TableCell className="text-center font-medium">{airport.country}</TableCell>
                      <TableCell className="text-center font-medium">{airport.terminals}</TableCell>
                      <TableCell className="text-center font-medium">{airport.runways}</TableCell>
                      <TableCell className="text-center">{getStatusBadge(airport.status)}</TableCell>
                      <TableCell>
                        <div className="flex justify-center gap-2">
                          <Dialog
                            open={editingAirport?.id === airport.id}
                            onOpenChange={(open) => {
                              if (!open) {
                                setEditingAirport(null)
                                setErrors({})
                              }
                            }}
                          >
                            <DialogTrigger asChild>
                              <Button
                                size="sm"
                                className="bg-cyan-500 hover:bg-cyan-600 text-white h-8 px-3"
                                onClick={() => handleEdit(airport)}
                              >
                                <Edit className="h-3 w-3 mr-1" />
                                Sửa
                              </Button>
                            </DialogTrigger>
                            <DialogContent className="max-w-2xl max-h-[90vh] overflow-y-auto">
                              <DialogHeader>
                                <DialogTitle>Chỉnh Sửa Sân Bay</DialogTitle>
                              </DialogHeader>
                              <div className="grid grid-cols-2 gap-4">
                                {renderFormField(
                                  "code",
                                  "Mã Sân Bay",
                                  newAirport.code,
                                  (value) => handleInputChange("code", value),
                                  "text",
                                  true,
                                  "HAN",
                                  null,
                                  3,
                                )}
                                {renderFormField(
                                  "iataCode",
                                  "Mã IATA",
                                  newAirport.iataCode,
                                  (value) => handleInputChange("iataCode", value),
                                  "text",
                                  true,
                                  "HAN",
                                  null,
                                  3,
                                )}
                                {renderFormField(
                                  "icaoCode",
                                  "Mã ICAO",
                                  newAirport.icaoCode,
                                  (value) => handleInputChange("icaoCode", value),
                                  "text",
                                  true,
                                  "VVNB",
                                  null,
                                  4,
                                )}
                                {renderFormField(
                                  "name",
                                  "Tên Sân Bay",
                                  newAirport.name,
                                  (value) => handleInputChange("name", value),
                                  "text",
                                  true,
                                  "Sân bay Quốc tế Nội Bài",
                                )}
                                {renderFormField(
                                  "city",
                                  "Thành Phố",
                                  newAirport.city,
                                  (value) => handleInputChange("city", value),
                                  "text",
                                  true,
                                  "Hà Nội",
                                )}
                                {renderFormField(
                                  "country",
                                  "Quốc Gia",
                                  newAirport.country,
                                  (value) => handleInputChange("country", value),
                                  "text",
                                  true,
                                  "Chọn quốc gia",
                                  countries,
                                )}
                                {renderFormField(
                                  "terminals",
                                  "Số Nhà Ga",
                                  newAirport.terminals,
                                  (value) => handleInputChange("terminals", value),
                                  "number",
                                  false,
                                  "2",
                                )}
                                {renderFormField(
                                  "runways",
                                  "Số Đường Băng",
                                  newAirport.runways,
                                  (value) => handleInputChange("runways", value),
                                  "number",
                                  false,
                                  "2",
                                )}
                                {renderFormField(
                                  "timezone",
                                  "Múi Giờ",
                                  newAirport.timezone,
                                  (value) => handleInputChange("timezone", value),
                                  "text",
                                  false,
                                  "Chọn múi giờ",
                                  timezones,
                                )}
                                {renderFormField(
                                  "status",
                                  "Trạng Thái",
                                  newAirport.status,
                                  (value) => handleInputChange("status", value),
                                  "text",
                                  false,
                                  "Chọn trạng thái",
                                  statuses,
                                )}
                              </div>
                              <DialogFooter className="mt-4">
                                <Button
                                  variant="outline"
                                  onClick={() => {
                                    setEditingAirport(null)
                                    setErrors({})
                                  }}
                                >
                                  Hủy
                                </Button>
                                <Button onClick={handleUpdate} className="bg-blue-600 hover:bg-blue-700">
                                  Cập Nhật
                                </Button>
                              </DialogFooter>
                            </DialogContent>
                          </Dialog>
                          <Button
                            size="sm"
                            variant="destructive"
                            onClick={() => confirmDelete(airport.id)}
                            className="bg-red-500 hover:bg-red-600 text-white h-8 px-3"
                          >
                            <Trash2 className="h-3 w-3 mr-1" />
                            Xóa
                          </Button>
                        </div>
                      </TableCell>
                    </TableRow>
                  ))}
                  {/* Thêm các dòng trống để duy trì chiều cao tối thiểu */}
                  {filteredAirports.length < 5 &&
                    Array.from({ length: 5 - filteredAirports.length }).map((_, index) => (
                      <TableRow key={`empty-${index}`} className="h-16">
                        <TableCell colSpan={11} className="text-center text-transparent">
                          -
                        </TableCell>
                      </TableRow>
                    ))}
                </TableBody>
              </Table>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Thông báo không tìm thấy - cố định vị trí */}
      <div className="h-16 flex items-center justify-center">
        {filteredAirports.length === 0 && searchQuery && (
          <div className="text-center text-gray-500">
            Không tìm thấy sân bay nào phù hợp với từ khóa tìm kiếm "{searchQuery}".
          </div>
        )}
      </div>

      {/* Dialog xác nhận xóa */}
      <AlertDialog open={isDeleteDialogOpen} onOpenChange={setIsDeleteDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Xác nhận xóa sân bay</AlertDialogTitle>
            <AlertDialogDescription>
              Bạn có chắc chắn muốn xóa sân bay này? Hành động này không thể hoàn tác.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel onClick={() => setIsDeleteDialogOpen(false)}>Hủy</AlertDialogCancel>
            <AlertDialogAction onClick={handleDelete} className="bg-red-500 hover:bg-red-600">
              Xóa
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}
