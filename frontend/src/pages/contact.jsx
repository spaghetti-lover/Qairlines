'use client'

import { useState } from 'react'
import { MapPin, Phone, Mail, Github, Linkedin, Facebook,  } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Alert, AlertDescription } from "@/components/ui/alert"

const TrangLienHe = () => {
  const [formData, setFormData] = useState({
    ten: '',
    email: '',
    chuDe: '',
    tinNhan: ''
  })
  const [dangGui, setDangGui] = useState(false)
  const [thongBaoGui, setThongBaoGui] = useState('')

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData(prevState => ({ ...prevState, [name]: value }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setDangGui(true)
    // Giả lập gửi form
    await new Promise(resolve => setTimeout(resolve, 1000))
    setThongBaoGui('Cảm ơn bạn đã gửi tin nhắn. Chúng tôi sẽ liên hệ lại sớm!')
    setDangGui(false)
    setFormData({ ten: '', email: '', chuDe: '', tinNhan: '' })
  }

  return (
    <div className="min-h-screen bg-background py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold text-center mb-12">Liên Hệ Với Chúng Tôi</h1>
        
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          {/* Thông tin liên hệ */}
          <Card className="space-y-4 border-orange border-2">
            <CardHeader>
              <CardTitle>Thông Tin Liên Hệ</CardTitle>
              <CardDescription>Hãy liên lạc với chúng tôi</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="flex items-center">
                  <MapPin className="h-6 w-6 text-primary mr-2" />
                  <span>Nhà E3, 144 Xuân Thủy, quận Cầu Giấy, Hà Nội , Việt Nam</span>
                </div>
                <div className="flex items-center">
                  <Phone className="h-6 w-6 text-primary mr-2" />
                  <span>+84 123 456 789</span>
                </div>
                <div className="flex items-center">
                  <Mail className="h-6 w-6 text-primary mr-2" />
                  <span>qairline-support@qairline.website</span>
                </div>
              </div>
              
              <h3 className="text-xl font-semibold mt-8 mb-4">Theo Dõi Chúng Tôi</h3>
              <div className="flex space-x-4">
                <a href="https://github.com/oceantran27/QAirline.git" target="_blank" rel="noopener noreferrer" className="text-muted-foreground hover:text-orange">
                  <span className=" sr-only">GitHub</span>
                  <Github className="h-6 w-6" />
                </a>
                <a href="https://www.linkedin.com/in/h%C6%B0ng-nguy%E1%BB%85n-duy-685477295/" target="_blank" rel="noopener noreferrer" className="text-muted-foreground hover:text-orange">
                  <span className="sr-only">LinkedIn</span>
                  <Linkedin className="h-6 w-6" />
                </a>
                <a href="https://www.facebook.com/profile.php?id=100045370126663" target="_blank" rel="noopener noreferrer" className="text-muted-foreground hover:text-orange">
                  <span className="sr-only">LinkedIn</span>
                  <Facebook className="h-6 w-6" />
                </a>
              </div>
            </CardContent>
          </Card>

          {/* Form liên hệ */}
          <Card className="space-y-4 border-orange border-2">
            <CardHeader>
              <CardTitle>Gửi Tin Nhắn Cho Chúng Tôi</CardTitle>
              <CardDescription>Điền vào mẫu dưới đây để liên hệ</CardDescription>
            </CardHeader>
            <CardContent>
              <form onSubmit={handleSubmit} className="space-y-6">
                <div className="space-y-2">
                  <Label htmlFor="ten">Tên</Label>
                  <Input
                    type="text"
                    name="ten"
                    id="ten"
                    value={formData.ten}
                    onChange={handleChange}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="email">Email</Label>
                  <Input
                    type="email"
                    name="email"
                    id="email"
                    value={formData.email}
                    onChange={handleChange}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="chuDe">Chủ đề</Label>
                  <Input
                    type="text"
                    name="chuDe"
                    id="chuDe"
                    value={formData.chuDe}
                    onChange={handleChange}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="tinNhan">Tin nhắn</Label>
                  <Textarea
                    name="tinNhan"
                    id="tinNhan"
                    rows={4}
                    value={formData.tinNhan}
                    onChange={handleChange}
                    required
                  />
                </div>
                <Button variant="orange" type="submit" disabled={dangGui}>
                  {dangGui ? 'Đang gửi...' : 'Gửi Tin Nhắn'}
                </Button>
              </form>
              {thongBaoGui && (
                <Alert className="mt-4">
                  <AlertDescription>{thongBaoGui}</AlertDescription>
                </Alert>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}

export default TrangLienHe

