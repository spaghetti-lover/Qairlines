"use client"

import { useState } from "react"
import { Search, Plus, Edit, Trash2, Plane, AlertCircle } from "lucide-react"
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

const hardcodedAircrafts = [
  {
    id: 1,
    code: "VN-A321",
    model: "A321-200",
    manufacturer: "Airbus",
    seats: 180,
    businessSeats: 12,
    economySeats: 168,
    yearManufactured: 2018,
    status: "Hoạt động",
    lastMaintenance: "2024-01-15",
    nextMaintenance: "2024-07-15",
  },
  {
    id: 2,
    code: "VN-B737",
    model: "B737-800",
    manufacturer: "Boeing",
    seats: 160,
    businessSeats: 8,
    economySeats: 152,
    yearManufactured: 2019,
    status: "Hoạt động",
    lastMaintenance: "2024-02-10",
    nextMaintenance: "2024-08-10",
  },
  {
    id: 3,
    code: "VN-A350",
    model: "A350-900",
    manufacturer: "Airbus",
    seats: 300,
    businessSeats: 28,
    economySeats: 272,
    yearManufactured: 2020,
    status: "Bảo trì",
    lastMaintenance: "2024-03-01",
    nextMaintenance: "2024-09-01",
  },
  {
    id: 4,
    code: "VN-B787",
    model: "B787-9",
    manufacturer: "Boeing",
    seats: 280,
    businessSeats: 24,
    economySeats: 256,
    yearManufactured: 2021,
    status: "Hoạt động",
    lastMaintenance: "2024-01-20",
    nextMaintenance: "2024-07-20",
  },
  {
    id: 5,
    code: "VN-A319",
    model: "A319-100",
    manufacturer: "Airbus",
    seats: 140,
    businessSeats: 8,
    economySeats: 132,
    yearManufactured: 2017,
    status: "Ngừng hoạt động",
    lastMaintenance: "2023-12-05",
    nextMaintenance: "2024-06-05",
  },
]

export default function AircraftManagement() {
  const [aircrafts, setAircrafts] = useState(hardcodedAircrafts)
  const [searchQuery, setSearchQuery] = useState("")
  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false)
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false)
  const [deleteAircraftId, setDeleteAircraftId] = useState(null)
  const [editingAircraft, setEditingAircraft] = useState(null)
  const [errors, setErrors] = useState({})
  const [newAircraft, setNewAircraft] = useState({
    code: "",
    model: "",
    manufacturer: "",
    seats: "",
    businessSeats: "",
    economySeats: "",
    yearManufactured: "",
    status: "Hoạt động",
  })

  const getStatusBadge = (status) => {
    switch (status) {
      case "Hoạt động":
        return <Badge className="bg-green-100 text-green-800 hover:bg-green-100">Hoạt động</Badge>
      case "Bảo trì":
        return <Badge className="bg-yellow-100 text-yellow-800 hover:bg-yellow-100">Bảo trì</Badge>
      case "Ngừng hoạt động":
        return <Badge className="bg-red-100 text-red-800 hover:bg-red-100">Ngừng hoạt động</Badge>
      default:
        return <Badge variant="secondary">{status}</Badge>
    }
  }

  const safeParseInt = (value) => {
    const parsed = Number.parseInt(value)
    return isNaN(parsed) ? 0 : parsed
  }

  const validateForm = (aircraft, isEdit = false) => {
    const newErrors = {}

    // Validate mã tàu bay
    if (!aircraft.code.trim()) {
      newErrors.code = "Mã tàu bay là bắt buộc"
    } else if (!/^[A-Z]{2}-[A-Z0-9]{3,6}$/i.test(aircraft.code.trim())) {
      newErrors.code = "Mã tàu bay phải có định dạng XX-XXXXX (VD: VN-A321)"
    } else {
      // Kiểm tra trùng lặp mã tàu bay
      const existingAircraft = aircrafts.find(
        (a) => a.code.toLowerCase() === aircraft.code.trim().toLowerCase() && (!isEdit || a.id !== editingAircraft?.id),
      )
      if (existingAircraft) {
        newErrors.code = "Mã tàu bay đã tồn tại"
      }
    }

    // Validate loại tàu bay
    if (!aircraft.model.trim()) {
      newErrors.model = "Loại tàu bay là bắt buộc"
    } else if (aircraft.model.trim().length < 2) {
      newErrors.model = "Loại tàu bay phải có ít nhất 2 ký tự"
    }

    // Validate hãng sản xuất
    if (!aircraft.manufacturer) {
      newErrors.manufacturer = "Hãng sản xuất là bắt buộc"
    }

    // Validate năm sản xuất - phải là số hợp lệ
    if (aircraft.yearManufactured && !/^\d{4}$/.test(aircraft.yearManufactured.toString().trim())) {
      newErrors.yearManufactured = "Năm sản xuất phải là số 4 chữ số"
    } else if (aircraft.yearManufactured) {
      const currentYear = new Date().getFullYear()
      const year = safeParseInt(aircraft.yearManufactured)
      if (year < 1950 || year > currentYear + 2) {
        newErrors.yearManufactured = `Năm sản xuất phải từ 1950 đến ${currentYear + 2}`
      }
    }

    // Validate số ghế - phải là số hợp lệ và lớn hơn 0
    if (!aircraft.seats || aircraft.seats.toString().trim() === "") {
      newErrors.seats = "Tổng số ghế là bắt buộc"
    } else if (!/^\d+$/.test(aircraft.seats.toString().trim())) {
      newErrors.seats = "Tổng số ghế phải là số nguyên dương"
    } else if (safeParseInt(aircraft.seats) <= 0) {
      newErrors.seats = "Tổng số ghế phải lớn hơn 0"
    } else if (safeParseInt(aircraft.seats) > 1000) {
      newErrors.seats = "Số ghế không được vượt quá 1000"
    }

    if (!aircraft.businessSeats || aircraft.businessSeats.toString().trim() === "") {
      newErrors.businessSeats = "Số ghế thương gia là bắt buộc"
    } else if (!/^\d+$/.test(aircraft.businessSeats.toString().trim())) {
      newErrors.businessSeats = "Số ghế thương gia phải là số nguyên"
    } else if (safeParseInt(aircraft.businessSeats) < 0) {
      newErrors.businessSeats = "Số ghế thương gia không được âm"
    }

    if (!aircraft.economySeats || aircraft.economySeats.toString().trim() === "") {
      newErrors.economySeats = "Số ghế phổ thông là bắt buộc"
    } else if (!/^\d+$/.test(aircraft.economySeats.toString().trim())) {
      newErrors.economySeats = "Số ghế phổ thông phải là số nguyên"
    } else if (safeParseInt(aircraft.economySeats) < 0) {
      newErrors.economySeats = "Số ghế phổ thông không được âm"
    }

    // Validate tổng ghế = ghế thương gia + ghế phổ thông (chỉ khi tất cả đều có giá trị hợp lệ)
    if (
      aircraft.seats &&
      aircraft.businessSeats &&
      aircraft.economySeats &&
      !newErrors.seats &&
      !newErrors.businessSeats &&
      !newErrors.economySeats
    ) {
      const totalSeats = safeParseInt(aircraft.seats)
      const businessSeats = safeParseInt(aircraft.businessSeats)
      const economySeats = safeParseInt(aircraft.economySeats)

      if (totalSeats !== businessSeats + economySeats) {
        newErrors.seats = "Tổng số ghế phải bằng ghế thương gia + ghế phổ thông"
        newErrors.businessSeats = "Tổng không khớp với tổng số ghế"
        newErrors.economySeats = "Tổng không khớp với tổng số ghế"
      }
    }

    return newErrors
  }

  const handleInputChange = (field, value) => {
    // Validation real-time cho các trường số - chỉ cho phép số và không cho phép số 0 ở đầu
    if (["seats", "businessSeats", "economySeats"].includes(field)) {
      // Chỉ cho phép số, không cho phép ký tự khác
      if (value && !/^\d*$/.test(value)) {
        return // Không cho phép nhập ký tự không phải số
      }
      // Không cho phép số 0 ở đầu (trừ khi chỉ có một số 0)
      if (value.length > 1 && value.startsWith("0")) {
        return
      }
    }

    if (field === "yearManufactured") {
      // Chỉ cho phép số và tối đa 4 chữ số
      if (value && (!/^\d*$/.test(value) || value.length > 4)) {
        return // Không cho phép nhập
      }
      // Không cho phép số 0 ở đầu
      if (value.length > 1 && value.startsWith("0")) {
        return
      }
    }

    setNewAircraft({ ...newAircraft, [field]: value })

    // Clear error khi user bắt đầu sửa
    if (errors[field]) {
      setErrors({ ...errors, [field]: null })
    }

    // Auto-calculate tổng ghế khi thay đổi ghế thương gia hoặc phổ thông
    if (field === "businessSeats" || field === "economySeats") {
      const businessSeats = field === "businessSeats" ? safeParseInt(value) : safeParseInt(newAircraft.businessSeats)
      const economySeats = field === "economySeats" ? safeParseInt(value) : safeParseInt(newAircraft.economySeats)

      if (businessSeats >= 0 && economySeats >= 0) {
        const totalCalculated = businessSeats + economySeats
        setNewAircraft((prev) => ({
          ...prev,
          [field]: value,
          seats: totalCalculated > 0 ? totalCalculated.toString() : "",
        }))
      }
    }
  }

  const handleAdd = () => {
    const validationErrors = validateForm(newAircraft)
    setErrors(validationErrors)

    if (Object.keys(validationErrors).length > 0) {
      toast({
        title: "Lỗi validation",
        description: "Vui lòng kiểm tra lại thông tin đã nhập.",
        variant: "destructive",
      })
      return
    }

    const maxId = aircrafts.length > 0 ? Math.max(...aircrafts.map((a) => a.id || 0)) : 0

    const aircraft = {
      id: maxId + 1,
      code: newAircraft.code.trim().toUpperCase(),
      model: newAircraft.model.trim(),
      manufacturer: newAircraft.manufacturer,
      seats: safeParseInt(newAircraft.seats),
      businessSeats: safeParseInt(newAircraft.businessSeats),
      economySeats: safeParseInt(newAircraft.economySeats),
      yearManufactured: safeParseInt(newAircraft.yearManufactured) || null,
      status: newAircraft.status,
      lastMaintenance: new Date().toISOString().split("T")[0],
      nextMaintenance: new Date(Date.now() + 180 * 24 * 60 * 60 * 1000).toISOString().split("T")[0],
    }

    setAircrafts([...aircrafts, aircraft])
    setNewAircraft({
      code: "",
      model: "",
      manufacturer: "",
      seats: "",
      businessSeats: "",
      economySeats: "",
      yearManufactured: "",
      status: "Hoạt động",
    })
    setErrors({})
    setIsAddDialogOpen(false)
    toast({
      title: "Thành công",
      description: "Thêm tàu bay mới thành công.",
    })
  }

  const handleEdit = (aircraft) => {
    setEditingAircraft(aircraft)
    setNewAircraft({
      code: aircraft.code || "",
      model: aircraft.model || "",
      manufacturer: aircraft.manufacturer || "",
      seats: (aircraft.seats || 0).toString(),
      businessSeats: (aircraft.businessSeats || 0).toString(),
      economySeats: (aircraft.economySeats || 0).toString(),
      yearManufactured: (aircraft.yearManufactured || "").toString(),
      status: aircraft.status || "Hoạt động",
    })
    setErrors({})
  }

  const handleUpdate = () => {
    const validationErrors = validateForm(newAircraft, true)
    setErrors(validationErrors)

    if (Object.keys(validationErrors).length > 0) {
      toast({
        title: "Lỗi validation",
        description: "Vui lòng kiểm tra lại thông tin đã nhập.",
        variant: "destructive",
      })
      return
    }

    const updatedAircrafts = aircrafts.map((aircraft) =>
      aircraft.id === editingAircraft.id
        ? {
            ...aircraft,
            code: newAircraft.code.trim().toUpperCase(),
            model: newAircraft.model.trim(),
            manufacturer: newAircraft.manufacturer,
            seats: safeParseInt(newAircraft.seats),
            businessSeats: safeParseInt(newAircraft.businessSeats),
            economySeats: safeParseInt(newAircraft.economySeats),
            yearManufactured: safeParseInt(newAircraft.yearManufactured) || null,
            status: newAircraft.status,
          }
        : aircraft,
    )

    setAircrafts(updatedAircrafts)
    setEditingAircraft(null)
    setNewAircraft({
      code: "",
      model: "",
      manufacturer: "",
      seats: "",
      businessSeats: "",
      economySeats: "",
      yearManufactured: "",
      status: "Hoạt động",
    })
    setErrors({})
    toast({
      title: "Thành công",
      description: "Cập nhật thông tin tàu bay thành công.",
    })
  }

  const confirmDelete = (id) => {
    setDeleteAircraftId(id)
    setIsDeleteDialogOpen(true)
  }

  const handleDelete = () => {
    setAircrafts(aircrafts.filter((aircraft) => aircraft.id !== deleteAircraftId))
    setIsDeleteDialogOpen(false)
    setDeleteAircraftId(null)
    toast({
      title: "Thành công",
      description: "Xóa tàu bay thành công.",
    })
  }

  const filteredAircrafts = aircrafts.filter(
    (aircraft) =>
      (aircraft.code || "").toLowerCase().includes(searchQuery.toLowerCase()) ||
      (aircraft.model || "").toLowerCase().includes(searchQuery.toLowerCase()) ||
      (aircraft.manufacturer || "").toLowerCase().includes(searchQuery.toLowerCase()) ||
      (aircraft.status || "").toLowerCase().includes(searchQuery.toLowerCase()),
  )

  const stats = {
    total: aircrafts.length,
    active: aircrafts.filter((a) => a.status === "Hoạt động").length,
    maintenance: aircrafts.filter((a) => a.status === "Bảo trì").length,
    inactive: aircrafts.filter((a) => a.status === "Ngừng hoạt động").length,
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

  const handleSearch = () => {
    // The search is already handled by the filteredAircrafts useMemo/filter
    // This function can be used for additional search logic if needed
    console.log("Searching for:", searchQuery)
  }

  return (
    <div className="pt-10 pl-64 mx-auto">
      {/* Header */}
      <div className="flex justify-between items-center mb-6">
        <div className="flex items-center gap-3">
          <Plane className="h-8 w-8 text-blue-600" />
          <h1 className="text-3xl font-bold text-gray-900">Quản Lý Tàu Bay</h1>
        </div>
        <Dialog open={isAddDialogOpen} onOpenChange={setIsAddDialogOpen}>
          <DialogTrigger asChild>
            <Button className="bg-blue-600 hover:bg-blue-700 text-white">
              <Plus className="h-4 w-4 mr-2" />
              Thêm Tàu Bay Mới
            </Button>
          </DialogTrigger>
          <DialogContent className="max-w-2xl max-h-[90vh] overflow-y-auto">
            <DialogHeader>
              <DialogTitle>Thêm Tàu Bay Mới</DialogTitle>
            </DialogHeader>
            <div className="grid grid-cols-2 gap-4">
              {renderFormField(
                "code",
                "Mã Tàu Bay",
                newAircraft.code,
                (value) => handleInputChange("code", value),
                "text",
                true,
                "VN-A321",
              )}
              {renderFormField(
                "model",
                "Loại",
                newAircraft.model,
                (value) => handleInputChange("model", value),
                "text",
                true,
                "A321-200",
              )}
              {renderFormField(
                "manufacturer",
                "Hãng Sản Xuất",
                newAircraft.manufacturer,
                (value) => handleInputChange("manufacturer", value),
                "text",
                true,
                "Chọn hãng sản xuất",
                [
                  { value: "Airbus", label: "Airbus" },
                  { value: "Boeing", label: "Boeing" },
                  { value: "Embraer", label: "Embraer" },
                  { value: "Bombardier", label: "Bombardier" },
                ],
              )}
              {renderFormField(
                "yearManufactured",
                "Năm Sản Xuất",
                newAircraft.yearManufactured,
                (value) => handleInputChange("yearManufactured", value),
                "number",
                false,
                "2020",
              )}
              {renderFormField(
                "businessSeats",
                "Ghế Thương Gia",
                newAircraft.businessSeats,
                (value) => handleInputChange("businessSeats", value),
                "number",
                true,
                "12",
              )}
              {renderFormField(
                "economySeats",
                "Ghế Phổ Thông",
                newAircraft.economySeats,
                (value) => handleInputChange("economySeats", value),
                "number",
                true,
                "168",
              )}
              {renderFormField(
                "seats",
                "Tổng Số Ghế",
                newAircraft.seats,
                (value) => handleInputChange("seats", value),
                "number",
                true,
                "180",
              )}
              {renderFormField(
                "status",
                "Trạng Thái",
                newAircraft.status,
                (value) => handleInputChange("status", value),
                "text",
                false,
                "",
                [
                  { value: "Hoạt động", label: "Hoạt động" },
                  { value: "Bảo trì", label: "Bảo trì" },
                  { value: "Ngừng hoạt động", label: "Ngừng hoạt động" },
                ],
              )}
            </div>
            <DialogFooter className="mt-4">
              <Button
                variant="outline"
                onClick={() => {
                  setIsAddDialogOpen(false)
                  setErrors({})
                  setNewAircraft({
                    code: "",
                    model: "",
                    manufacturer: "",
                    seats: "",
                    businessSeats: "",
                    economySeats: "",
                    yearManufactured: "",
                    status: "Hoạt động",
                  })
                }}
              >
                Hủy
              </Button>
              <Button onClick={handleAdd} className="bg-blue-600 hover:bg-blue-700">
                Thêm Tàu Bay
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium text-gray-600">Tổng Số Tàu Bay</CardTitle>
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
            <CardTitle className="text-sm font-medium text-gray-600">Ngừng Hoạt Động</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-red-600">{stats.inactive}</div>
          </CardContent>
        </Card>
      </div>

      {/* Search */}
      <div className="relative mb-6 flex">
        <Input
          type="text"
          placeholder="Tìm kiếm theo mã tàu bay, loại, hãng sản xuất hoặc trạng thái"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onKeyPress={(e) => e.key === "Enter" && handleSearch()}
          className="h-12 rounded-r-none border-r-0 focus:border-r-0"
        />
        <Button onClick={handleSearch} className="h-12 px-4 bg-blue-600 hover:bg-blue-700 rounded-l-none">
          <Search className="h-4 w-4" />
        </Button>
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
                    <TableHead className="text-center font-semibold">Mã Tàu Bay</TableHead>
                    <TableHead className="text-center font-semibold">Loại</TableHead>
                    <TableHead className="text-center font-semibold">Hãng SX</TableHead>
                    <TableHead className="text-center font-semibold">Năm SX</TableHead>
                    <TableHead className="text-center font-semibold">Tổng Ghế</TableHead>
                    <TableHead className="text-center font-semibold">Thương Gia</TableHead>
                    <TableHead className="text-center font-semibold">Phổ Thông</TableHead>
                    <TableHead className="text-center font-semibold">Trạng Thái</TableHead>
                    <TableHead className="text-center font-semibold">Thao Tác</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {filteredAircrafts.map((aircraft, index) => (
                    <TableRow key={aircraft.id} className={index % 2 === 0 ? "bg-white" : "bg-gray-50"}>
                      <TableCell className="text-center font-medium">{index + 1}</TableCell>
                      <TableCell className="text-center font-semibold text-blue-600">{aircraft.code || "-"}</TableCell>
                      <TableCell className="text-center font-medium">{aircraft.model || "-"}</TableCell>
                      <TableCell className="text-center font-medium">{aircraft.manufacturer || "-"}</TableCell>
                      <TableCell className="text-center font-medium">{aircraft.yearManufactured || "-"}</TableCell>
                      <TableCell className="text-center font-medium">{aircraft.seats || 0}</TableCell>
                      <TableCell className="text-center font-medium">{aircraft.businessSeats || 0}</TableCell>
                      <TableCell className="text-center font-medium">{aircraft.economySeats || 0}</TableCell>
                      <TableCell className="text-center">{getStatusBadge(aircraft.status)}</TableCell>
                      <TableCell>
                        <div className="flex justify-center gap-2">
                          <Dialog
                            open={editingAircraft?.id === aircraft.id}
                            onOpenChange={(open) => {
                              if (!open) {
                                setEditingAircraft(null)
                                setErrors({})
                              }
                            }}
                          >
                            <DialogTrigger asChild>
                              <Button
                                size="sm"
                                className="bg-cyan-500 hover:bg-cyan-600 text-white h-8 px-3"
                                onClick={() => handleEdit(aircraft)}
                              >
                                <Edit className="h-3 w-3 mr-1" />
                                Sửa
                              </Button>
                            </DialogTrigger>
                            <DialogContent className="max-w-2xl max-h-[90vh] overflow-y-auto">
                              <DialogHeader>
                                <DialogTitle>Chỉnh Sửa Tàu Bay</DialogTitle>
                              </DialogHeader>
                              <div className="grid grid-cols-2 gap-4">
                                {renderFormField(
                                  "code",
                                  "Mã Tàu Bay",
                                  newAircraft.code,
                                  (value) => handleInputChange("code", value),
                                  "text",
                                  true,
                                  "VN-A321",
                                )}
                                {renderFormField(
                                  "model",
                                  "Loại",
                                  newAircraft.model,
                                  (value) => handleInputChange("model", value),
                                  "text",
                                  true,
                                  "A321-200",
                                )}
                                {renderFormField(
                                  "manufacturer",
                                  "Hãng Sản Xuất",
                                  newAircraft.manufacturer,
                                  (value) => handleInputChange("manufacturer", value),
                                  "text",
                                  true,
                                  "Chọn hãng sản xuất",
                                  [
                                    { value: "Airbus", label: "Airbus" },
                                    { value: "Boeing", label: "Boeing" },
                                    { value: "Embraer", label: "Embraer" },
                                    { value: "Bombardier", label: "Bombardier" },
                                  ],
                                )}
                                {renderFormField(
                                  "yearManufactured",
                                  "Năm Sản Xuất",
                                  newAircraft.yearManufactured,
                                  (value) => handleInputChange("yearManufactured", value),
                                  "number",
                                  false,
                                  "2020",
                                )}
                                {renderFormField(
                                  "businessSeats",
                                  "Ghế Thương Gia",
                                  newAircraft.businessSeats,
                                  (value) => handleInputChange("businessSeats", value),
                                  "number",
                                  true,
                                  "12",
                                )}
                                {renderFormField(
                                  "economySeats",
                                  "Ghế Phổ Thông",
                                  newAircraft.economySeats,
                                  (value) => handleInputChange("economySeats", value),
                                  "number",
                                  true,
                                  "168",
                                )}
                                {renderFormField(
                                  "seats",
                                  "Tổng Số Ghế",
                                  newAircraft.seats,
                                  (value) => handleInputChange("seats", value),
                                  "number",
                                  true,
                                  "180",
                                )}
                                {renderFormField(
                                  "status",
                                  "Trạng Thái",
                                  newAircraft.status,
                                  (value) => handleInputChange("status", value),
                                  "text",
                                  false,
                                  "",
                                  [
                                    { value: "Hoạt động", label: "Hoạt động" },
                                    { value: "Bảo trì", label: "Bảo trì" },
                                    { value: "Ngừng hoạt động", label: "Ngừng hoạt động" },
                                  ],
                                )}
                              </div>
                              <DialogFooter className="mt-4">
                                <Button
                                  variant="outline"
                                  onClick={() => {
                                    setEditingAircraft(null)
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
                            onClick={() => confirmDelete(aircraft.id)}
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
                  {filteredAircrafts.length < 5 &&
                    Array.from({ length: 5 - filteredAircrafts.length }).map((_, index) => (
                      <TableRow key={`empty-${index}`} className="h-16">
                        <TableCell colSpan={10} className="text-center text-transparent">
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
        {filteredAircrafts.length === 0 && searchQuery && (
          <div className="text-center text-gray-500">
            Không tìm thấy tàu bay nào phù hợp với từ khóa tìm kiếm "{searchQuery}".
          </div>
        )}
      </div>

      {/* Dialog xác nhận xóa */}
      <AlertDialog open={isDeleteDialogOpen} onOpenChange={setIsDeleteDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Xác nhận xóa tàu bay</AlertDialogTitle>
            <AlertDialogDescription>
              Bạn có chắc chắn muốn xóa tàu bay này? Hành động này không thể hoàn tác.
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
