"use client"

import { useEffect, useState } from "react"
import { Search } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { useRouter } from "next/router"
import { toast } from "@/hooks/use-toast"
import { EditCustomerDialog } from "@/components/admin/EditCustomerDialog"

export default function CustomerManagement() {
    const router = useRouter()
        
    useEffect(() => {
        const token = localStorage.getItem('token')
        if (!token) {
            router.push('/admin')
        }
        else getAllCustomers()

    }, [router])    

    const [customers, setCustomers] = useState([])
    const [searchTerm, setSearchTerm] = useState("")
    const [selectedCustomer, setSelectedCustomer] = useState(null)
    const [editingCustomer, setEditingCustomer] = useState(null)

    const filteredCustomers = customers.filter(
        (customer) =>
        (`${customer.lastName} ${customer.firstName}`).toLowerCase().includes(searchTerm.toLowerCase()) ||
        customer.email.toLowerCase().includes(searchTerm.toLowerCase())
    )

    const getAllCustomers = async () => {
        const getAllCustomersApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/customer/all`
    
        try { 
            const response = await fetch(getAllCustomersApi, {
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
            setCustomers(res.data.map(a => ({
                "uid": a.uid,
                "firstName": a.firstName,
                "lastName": a.lastName,
                "email": a.email,
                "dateOfBirth": a.dateOfBirth.seconds? new Date(a.dateOfBirth.seconds*1000).toISOString().split('T')[0] : a.dateOfBirth.split('T')[0],
                "gender": a.gender,
                "role": "customer",
                "loyaltyPoints": a.loyaltyPoints,
                "createdAt": a.createdAt.seconds ? new Date(a.createdAt.seconds*1000).toISOString().split('T')[0] : a.createdAt.split('T')[0],
            })))
        } catch (error) {
          toast({
            title: "Lỗi",
            description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
            variant: "destructive"
          })
        }
    }

    const handleEdit = (customer) => {
      setEditingCustomer(customer)
    }

    const handleEditComplete = async (updatedCustomer) => {
        // Here you would typically make an API call to update the customer
        // For now, we'll just update the local state
        setCustomers(customers.map(c => c.uid === updatedCustomer.uid ? updatedCustomer : c))
        setEditingCustomer(null)
        toast({
            title: "Thành công",
            description: "Thông tin khách hàng đã được cập nhật",
            variant: "default"
        })
    }

    const handleDelete = async (customer) => {
      setCustomers(customers.filter(a => a.uid !== customer.uid))
      const deleteCustomerApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/customer/delete?id=${customer.uid}`

      try {
        const response = await fetch(deleteCustomerApi, 
          {
            method: "DELETE",
            headers: {
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            }, 
          })
        if (!response.ok) {
            throw new Error("Send request failed")
        }
        toast({
          title: "Thành công",
          description: "Tài khoản của khách hàng đã được xóa",
        })
      } catch (error) {
        toast({
          title: "Xóa tài khoản thất bại",
          description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
          variant: "destructive"
        })
      }
    }

  return (
    <div className="container mx-auto py-10 pl-64">
      <h1 className="text-2xl font-semibold mb-5">Quản lý khách hàng</h1>
      <div className="flex items-center space-x-2 mb-4">
        <Input
          placeholder="Tìm kiếm khách hàng..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="max-w-sm"
        />
        <Button type="submit" className="bg-blue-500 hover:bg-blue-600">
          <Search className="h-4 w-4 mr-2" />
          Tìm kiếm
        </Button>
      </div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Họ và tên</TableHead>
            <TableHead>Email</TableHead>
            <TableHead>Ngày sinh</TableHead>
            <TableHead>Giới tính</TableHead>
            <TableHead>Điểm tích lũy</TableHead>
            <TableHead>Thao tác</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {filteredCustomers.map((customer) => (
            <TableRow key={customer.uid}>
              <TableCell>{`${customer.lastName} ${customer.firstName}`}</TableCell>
              <TableCell>{customer.email}</TableCell>
              <TableCell>{customer.dateOfBirth}</TableCell>
              <TableCell>{customer.gender === "male" ? "Nam" : "Nữ"}</TableCell>
              <TableCell>{customer.loyaltyPoints}</TableCell>
              <TableCell>
                <div className="flex space-x-2">
                  <Dialog>
                    <DialogTrigger asChild>
                      <Button
                        variant="outline"
                        onClick={() => setSelectedCustomer(customer)}
                      >
                        Chi tiết
                      </Button>
                    </DialogTrigger>
                    <DialogContent className="sm:max-w-[425px]">
                      <DialogHeader>
                        <DialogTitle>Thông tin chi tiết khách hàng</DialogTitle>
                      </DialogHeader>
                      {selectedCustomer && (
                        <div className="grid gap-4 py-4">
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Họ và tên:</span>
                            <span className="col-span-3">{`${selectedCustomer.lastName} ${selectedCustomer.firstName}`}</span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Email:</span>
                            <span className="col-span-3">{selectedCustomer.email}</span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Ngày sinh:</span>
                            <span className="col-span-3">
                              {selectedCustomer.dateOfBirth}
                            </span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Giới tính:</span>
                            <span className="col-span-3">
                              {selectedCustomer.gender === "male" ? "Nam" : "Nữ"}
                            </span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Tích lũy:</span>
                            <span className="col-span-3">{selectedCustomer.loyaltyPoints}</span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Địa chỉ:</span>
                            <span className="col-span-3">{selectedCustomer.address || ''}</span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Số hộ chiếu:</span>
                            <span className="col-span-3">{selectedCustomer.passportNumber || ''}</span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">CMND/CCCD:</span>
                            <span className="col-span-3">{selectedCustomer.identificationNumber || ''}</span>
                          </div>
                          <div className="grid grid-cols-4 items-center gap-4">
                            <span className="font-bold">Ngày tạo:</span>
                            <span className="col-span-3">
                              {selectedCustomer.createdAt}
                            </span>
                          </div>
                        </div>
                      )}
                    </DialogContent>
                  </Dialog>
                  <Button
                    variant="outline"
                    onClick={() => handleEdit(customer)}
                    className="bg-yellow-500 hover:bg-yellow-600 text-white"
                  >
                    Sửa
                  </Button>
                  <Button
                    variant="outline"
                    onClick={() => handleDelete(customer)}
                    className="bg-red-500 hover:bg-red-600 text-white"
                  >
                    Xóa
                  </Button>
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      {editingCustomer && (
        <EditCustomerDialog
          customer={editingCustomer}
          onClose={() => setEditingCustomer(null)}
          onSave={handleEditComplete}
        />
      )}
    </div>
  )
}