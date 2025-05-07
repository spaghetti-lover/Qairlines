import React from 'react'
import { useState } from 'react'
import { Eye, EyeOff } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from "@/components/ui/dialog"
import { useRouter } from 'next/router'

export default function LoginForm() {
    const [showPassword, setShowPassword] = useState(false)
    const [formData, setFormData] = useState({
      email: '',
      password: ''
    })
    const [isErrorDialogOpen, setIsErrorDialogOpen] = useState(false)
    const router = useRouter()
  
    const handleInputChange = (e) => {
      const { name, value } = e.target
      setFormData(prevData => ({
        ...prevData,
        [name]: value
      }))
    }
  
    const handleSubmit = async (e) => {
        e.preventDefault()
        const loginApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/login/admin`
  
        try {
            const response = await fetch(loginApi, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                }, 
                body: JSON.stringify(formData)
            })
            if (!response.ok) {
                throw new Error("Login failed")
            }
  
            const data = await response.json()
            localStorage.setItem("token", data.token)
            router.push("/admin/dashboard")
        } catch (error) {
            console.error("Login error:", error)
            setIsErrorDialogOpen(true)
        }
    }
  
    return (
        <>
        <div className="flex items-center justify-center min-h-screen bg-cover bg-center"
            style={{
                backgroundImage: "url('/clouds-background.jpg')",
            }}
        >
        <Card className="w-full max-w-md bg-white shadow-xl">
            <CardHeader className="space-y-1">
            <CardTitle className="text-3xl font-bold text-center text-[#e8604c]">Đăng nhập admin</CardTitle>
        </CardHeader>
        <CardContent>
        <form onSubmit={handleSubmit} className="space-y-4">
            <div className="space-y-2">
                <Label htmlFor="email" className="text-sm font-medium text-gray-700">Email</Label>
                <Input
                    id="email"
                    name="email"
                    type="email"
                    placeholder="m@example.com"
                    required
                    className="w-full px-3 py-2 border rounded-md focus:ring-2 focus:ring-[#e8604c]"
                    value={formData.email}
                    onChange={handleInputChange}
                />
            </div>
            <div className="space-y-2">
                <Label htmlFor="password" className="text-sm font-medium text-gray-700">Mật khẩu</Label>
                <div className="relative">
                    <Input
                    id="password"
                    name="password"
                    type={showPassword ? "text" : "password"}
                    required
                    className="w-full px-3 py-2 border rounded-md focus:ring-2 focus:ring-[#e8604c]"
                    value={formData.password}
                    onChange={handleInputChange}
                    />
                    <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-400 hover:text-gray-600"
                    >
                    {showPassword ? <EyeOff className="h-5 w-5" /> : <Eye className="h-5 w-5" />}
                    </button>
                </div>
            </div>
            <div className="relative my-4">
                <div className="absolute inset-0 flex items-center">
                <span className="w-full border-t" />
                </div>
            </div>
            <Button type="submit" className="w-full bg-[#e8604c] hover:bg-[#d55643] text-white">
            Đăng nhập
            </Button>
        </form>
        </CardContent>
        </Card>
        </div>

        <Dialog open={isErrorDialogOpen} onOpenChange={setIsErrorDialogOpen}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Đăng nhập thất bại</DialogTitle>
                    <DialogDescription>Email hoặc mật khẩu không đúng. Vui lòng thử lại.</DialogDescription>
                </DialogHeader>
                <Button onClick={() => setIsErrorDialogOpen(false)}>Đóng</Button>
            </DialogContent>
        </Dialog>
    </>
    )
}

