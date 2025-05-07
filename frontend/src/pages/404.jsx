import React from 'react'
import Link from 'next/link'
import { AlertTriangle } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

export default function Custom404() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-cover bg-center"
         style={{
           backgroundImage: "url('/clouds-background.jpg')",
         }}
    >
      <Card className="w-full max-w-md bg-white shadow-xl">
        <CardHeader className="space-y-1">
          <CardTitle className="text-3xl font-bold text-center text-[#e8604c]">404 - Không Tìm Thấy Trang</CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col items-center space-y-4">
          <AlertTriangle className="w-20 h-20 text-[#e8604c]" aria-hidden="true" />
          <p className="text-center text-gray-600">
            Rất tiếc! Trang bạn đang tìm kiếm không tồn tại hoặc đã được di chuyển.
          </p>
          <Link href="/" passHref>
            <Button className="bg-[#e8604c] hover:bg-[#d55643] text-white">
              Quay Về Trang Chủ
            </Button>
          </Link>
        </CardContent>
      </Card>
    </div>
  )
}

