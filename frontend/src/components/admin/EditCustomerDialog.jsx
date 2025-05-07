import { useState } from 'react'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { toast } from '@/hooks/use-toast'

export function EditCustomerDialog({ customer, onClose, onSave }) {
  const [editedCustomer, setEditedCustomer] = useState(customer)

  const handleInputChange = (e) => {
    const { name, value } = e.target
    setEditedCustomer(prev => ({ ...prev, [name]: value }))
  }

  const handleSelectChange = (name, value) => {
    setEditedCustomer(prev => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const updateCustomerApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/customer/update?id=${editedCustomer.uid}`

    try {
        const response = await fetch(updateCustomerApi, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            }, 
            body: JSON.stringify(editedCustomer)
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }
        toast({
          title: "Thành công",
          description: "Thông tin của khách hàng đã được cập nhật",
        })
    } catch (error) {
      toast({
        title: "Cập nhật thông tin thất bại",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
    onSave(editedCustomer)
  }

  return (
    <Dialog open={true} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Chỉnh sửa thông tin khách hàng</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="firstName" className="text-right">
                Tên
              </Label>
              <Input
                id="firstName"
                name="firstName"
                value={editedCustomer.firstName}
                onChange={handleInputChange}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="lastName" className="text-right">
                Họ
              </Label>
              <Input
                id="lastName"
                name="lastName"
                value={editedCustomer.lastName}
                onChange={handleInputChange}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="phoneNumber" className="text-right">
                Số điện thoại
              </Label>
              <Input
                id="phoneNumber"
                name="phoneNumber"
                value={editedCustomer.phoneNumber || ''}
                onChange={handleInputChange}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="gender" className="text-right">
                Giới tính
              </Label>
              <Select name="gender" onValueChange={(value) => handleSelectChange('gender', value)} defaultValue={editedCustomer.gender}>
                <SelectTrigger className="col-span-3">
                  <SelectValue placeholder="Chọn giới tính" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="male">Nam</SelectItem>
                  <SelectItem value="female">Nữ</SelectItem>
                </SelectContent>
              </Select>
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="address" className="text-right">
                Địa chỉ
              </Label>
              <Input
                id="address"
                name="address"
                value={editedCustomer.address || ''}
                onChange={handleInputChange}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="passportNumber" className="text-right">
                Số hộ chiếu
              </Label>
              <Input
                id="passportNumber"
                name="passportNumber"
                value={editedCustomer.passportNumber || ''}
                onChange={handleInputChange}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="identificationNumber" className="text-right">
                Số CMND/CCCD
              </Label>
              <Input
                id="identificationNumber"
                name="identificationNumber"
                value={editedCustomer.identificationNumber || ''}
                onChange={handleInputChange}
                className="col-span-3"
              />
            </div>
          </div>
          <DialogFooter>
            <Button type="submit">Lưu thay đổi</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}