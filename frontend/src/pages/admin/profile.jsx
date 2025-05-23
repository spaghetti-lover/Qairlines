'use client'

import { useState, useEffect } from 'react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from "@/components/ui/alert-dialog"
import { useRouter } from 'next/router'
import { toast } from '@/hooks/use-toast'

export default function AdminProfilePage() {
  const router = useRouter()

  const [admin, setAdmin] = useState({
    uid: '',
    firstName: '',
    lastName: '',
    email: ''
  })
  const [editForm, setEditForm] = useState({ ...admin })
  const [oldPassword, setOldPassword] = useState('')
  const [newPassword, setNewPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) {
      router.push('/admin')
    }
    else getAdmin()

  }, [router])

  const getAdmin = async () => {
    const getAdminApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/admin`

    try {
        const response = await fetch(getAdminApi, {
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
        setAdmin({"uid": res.data.uid, "firstName": res.data.firstName, "lastName": res.data.lastName, "email": res.data.email})
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }

  const handleUpdateAdmin = async (e) => {
    e.preventDefault()
    const updateAdminApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/admin/update`

    try {
        const response = await fetch(updateAdminApi, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            },
            body: JSON.stringify(editForm)
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }
        toast({
          title: "Thành công",
          description: "Thông tin của bạn đã được cập nhật",
        })
        getAdmin()
    } catch (error) {
      toast({
        title: "Cập nhật thông tin thất bại",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }

  const handleDeleteAccount = async () => {
    const deleteAdminApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/admin/delete`

    try {
        const response = await fetch(deleteAdminApi, {
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
        title: "Xóa tài khoản thất bại",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }

    localStorage.removeItem('token')
    router.push('/admin')
  }

  const handlePasswordReset = async (e) => {
    e.preventDefault()
    const changePasswordApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/customer/change-password?`

    if (newPassword !== confirmPassword) {
      toast({
        title: "Lỗi",
        description: "Mật khẩu được gõ lại không chính xác!",
        variant: "destructive"
      })
      return
    }

    try {
        const response = await fetch(changePasswordApi +
          new URLSearchParams({
            id: admin.uid,
            admin: "true",
          }).toString(), {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "admin": "true",
                "authorization": "Bearer " + localStorage.getItem("token")
            },
            body: JSON.stringify({"email": admin.email, "oldPassword": oldPassword, "newPassword": newPassword})
        })
        if (!response.ok) {
            throw new Error("Send request failed")
        }
        toast({
          title: "Thành công",
          description: "Đổi mật khẩu thành công. Vui lòng đăng nhập lại",
        })
        router.push("/admin")
    } catch (error) {
      toast({
        title: "Đổi mật khẩu thất bại",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    router.push("/admin")
  }

  return (
    <div className="container mx-auto pt-10 pl-64 space-y-6">
      <h1 className="text-2xl font-semibold">Hồ Sơ Cá Nhân</h1>

      <Card>
        <CardHeader>
          <CardTitle>Thông tin cơ bản</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-2">
            <p><strong>UID:</strong> {admin.uid}</p>
            <p><strong>Tên:</strong> {admin.firstName}</p>
            <p><strong>Họ:</strong> {admin.lastName}</p>
            <p><strong>Email:</strong> {admin.email}</p>
          </div>
        </CardContent>
        <CardFooter>
          <Dialog>
            <DialogTrigger asChild>
              <Button>Chỉnh sửa</Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Thông tin hồ sơ</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleUpdateAdmin}>
                <div className="space-y-4">
                  <div>
                    <Label htmlFor="firstName">Tên</Label>
                    <Input
                      id="firstName"
                      value={editForm.firstName}
                      onChange={(e) => setEditForm({...editForm, firstName: e.target.value})}
                    />
                  </div>
                  <div>
                    <Label htmlFor="lastName">Họ</Label>
                    <Input
                      id="lastName"
                      value={editForm.lastName}
                      onChange={(e) => setEditForm({...editForm, lastName: e.target.value})}
                    />
                  </div>
                </div>
                <DialogFooter className="mt-4">
                  <Button type="submit">Lưu thay đổi</Button>
                </DialogFooter>
              </form>
            </DialogContent>
          </Dialog>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Đặt lại mật khẩu</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handlePasswordReset} className="space-y-4">
            <div>
              <Label htmlFor="oldPassword">Mật khẩu hiện tại</Label>
              <Input
                id="oldPassword"
                type="password"
                value={oldPassword}
                onChange={(e) => setOldPassword(e.target.value)}
              />
            </div>
            <div>
              <Label htmlFor="newPassword">Mật khẩu mới</Label>
              <Input
                id="newPassword"
                type="password"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
              />
            </div>
            <div>
              <Label htmlFor="confirmPassword">Gõ lại mật khẩu mới</Label>
              <Input
                id="confirmPassword"
                type="password"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
              />
            </div>
            <Button type="submit">Lưu thay đổi</Button>
          </form>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Cài đặt tài khoản</CardTitle>
        </CardHeader>
        <CardContent>
          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button variant="destructive">Xóa tài khoản</Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>Bạn có chắc chắn muốn xóa tài khoản?</AlertDialogTitle>
                <AlertDialogDescription>
                  Sẽ không thể khôi phục tài khoản một khi đã được xóa. Tất cả những dữ liệu liên quan tới tài khoản của bạn đều sẽ không được lưu lại.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel>Thoát</AlertDialogCancel>
                <AlertDialogAction onClick={handleDeleteAccount}>Xóa tài khoản</AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>

          <Button onClick={handleLogout} variant="outline" className="ml-4">Đăng xuất</Button>
        </CardContent>
      </Card>
    </div>
  )
}

