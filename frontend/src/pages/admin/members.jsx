'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'

import { Button } from "@/components/ui/button"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow, } from "@/components/ui/table"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger, } from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Plus } from 'lucide-react'
import { toast } from '@/hooks/use-toast'

export default function AdminManagementPage() {
  const router = useRouter()

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) {
      router.push('/admin')
    }
    else getAllAdmins()

  }, [router])

  const [admins, setAdmins] = useState([])
  const [isDialogOpen, setIsDialogOpen] = useState(false)
  const [errors, setErrors] = useState({})
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    repeatPassword: ''
  })

  const registerAdmin = async () => {
    const registerAdminApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/admin`

    try {
        const response = await fetch(registerAdminApi, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            },
            body: JSON.stringify(formData)
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }

        getAllAdmins()
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }

  const getAllAdmins = async () => {
    const getAllAdminsApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/admin/all`

    try {
        const response = await fetch(getAllAdminsApi, {
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
        setAdmins(res.data.map(a => {return {
          "name": `${a.firstName} ${a.lastName}`,
          "email": a.email,
          "createdAt": a.createdAt ? new Date(a.createdAt.seconds*1000).toISOString().split('T')[0] : '2024-12-21'
        }}))
    } catch (error) {
      console.log(error)
      toast({
        title: "Lỗi",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }

  const validateForm = () => {
    const newErrors = {}

    if (!formData.firstName.trim()) {
      newErrors.firstName = 'Thiếu tên'
    }

    if (!formData.lastName.trim()) {
      newErrors.lastName = 'Thiếu họ'
    }

    if (!formData.email.trim()) {
      newErrors.email = 'Thiếu email'
    } else if (!/\S+@\S+\.\S+/.test(formData.email)) {
      newErrors.email = 'Email không hợp lệ'
    }

    if (!formData.password) {
      newErrors.password = 'Thiếu mật khẩu'
    } else if (formData.password.length < 6) {
      newErrors.password = 'Độ dài mật khẩu tối thiểu 6 ký tự'
    }

    if (formData.password !== formData.repeatPassword) {
      newErrors.repeatPassword = 'Mật khẩu không khớp'
    }

    setErrors(newErrors)
    return Object.keys(newErrors).length === 0
  }

  const handleSubmit = (e) => {
    e.preventDefault()

    if (validateForm()) {
      // Send the data to backend
      registerAdmin()

      // Clear form after successful submission
      setFormData({
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        repeatPassword: ''
      })

      setIsDialogOpen(false)
      toast({
        title: "Thành công",
        description: "Quản trị viên mới đã được thêm vào danh sách",
      })
    }
  }

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: value
    }))
    // Clear error when user starts typing
    if (errors[name]) {
      setErrors(prev => ({
        ...prev,
        [name]: ''
      }))
    }
  }

  return (
    <div className="container mx-auto pt-10 pl-64">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-semibold">Quản Lý Admin</h1>
        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
          <DialogTrigger asChild>
            <Button className="bg-yellow-400 hover:bg-yellow-500 text-black font-medium">
              <Plus className="mr-2 h-4 w-4" />
              ADMIN MỚI
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Thêm admin mới</DialogTitle>
            </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4 pt-4">
              <div>
                <Input
                  type="text"
                  name="firstName"
                  placeholder="Tên"
                  value={formData.firstName}
                  onChange={handleChange}
                  className={errors.firstName ? 'border-red-500' : ''}
                />
                {errors.firstName && (
                  <p className="text-red-500 text-sm mt-1">{errors.firstName}</p>
                )}
              </div>

              <div>
                <Input
                  type="text"
                  name="lastName"
                  placeholder="Họ"
                  value={formData.lastName}
                  onChange={handleChange}
                  className={errors.lastName ? 'border-red-500' : ''}
                />
                {errors.lastName && (
                  <p className="text-red-500 text-sm mt-1">{errors.lastName}</p>
                )}
              </div>

              <div>
                <Input
                  type="email"
                  name="email"
                  placeholder="Email"
                  value={formData.email}
                  onChange={handleChange}
                  className={errors.email ? 'border-red-500' : ''}
                />
                {errors.email && (
                  <p className="text-red-500 text-sm mt-1">{errors.email}</p>
                )}
              </div>

              <div>
                <Input
                  type="password"
                  name="password"
                  placeholder="Mật Khẩu"
                  value={formData.password}
                  onChange={handleChange}
                  className={errors.password ? 'border-red-500' : ''}
                />
                {errors.password && (
                  <p className="text-red-500 text-sm mt-1">{errors.password}</p>
                )}
              </div>

              <div>
                <Input
                  type="password"
                  name="repeatPassword"
                  placeholder="Nhập Lại Mật Khẩu"
                  value={formData.repeatPassword}
                  onChange={handleChange}
                  className={errors.repeatPassword ? 'border-red-500' : ''}
                />
                {errors.repeatPassword && (
                  <p className="text-red-500 text-sm mt-1">{errors.repeatPassword}</p>
                )}
              </div>

              <Button
                type="submit"
                className="w-full bg-blue-500 hover:bg-blue-600"
              >
                Lưu
              </Button>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Tên</TableHead>
            <TableHead>Email</TableHead>
            <TableHead>Ngày đăng ký</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {admins.map((admin) => (
            <TableRow key={admin.uid}>
              <TableCell>{admin.name}</TableCell>
              <TableCell>{admin.email}</TableCell>
              <TableCell>{admin.createdAt}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  )
}

