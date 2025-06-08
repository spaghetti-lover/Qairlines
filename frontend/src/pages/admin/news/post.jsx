'use client'

import React, { useState, useCallback, useEffect } from 'react'
import Image from 'next/image'
import { useDropzone } from 'react-dropzone'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { useRouter } from 'next/router'
import { toast } from '@/hooks/use-toast'
import mockNewsDataService from '@/pages/mockNewsData'

export default function NewsPostingPage() {
  const router = useRouter()
  const { id } = router.query

  const [formData, setFormData] = useState({
    title: '',
    description: '',
    content: '',
    image: null,
    authorId: '1',
    authorName: 'Admin User'
  })
  const [previewImage, setPreviewImage] = useState(null)
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    if (id) {
      loadNewsData()
    }
    loadAuthorData()
  }, [id])

  const loadNewsData = async () => {
    try {
      const res = await mockNewsDataService.getNewsById(id)
      const newsData = res.data
      setFormData({
        title: newsData.title,
        description: newsData.description,
        content: newsData.content,
        image: newsData.image,
        authorId: newsData.authorId,
        authorName: newsData.authorName
      })
      setPreviewImage(newsData.image)
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Không thể tải thông tin bài viết",
        variant: "destructive"
      })
    }
  }

  const loadAuthorData = async () => {
    try {
      const res = await mockNewsDataService.getCurrentUser()
      setFormData(prev => ({
        ...prev,
        authorId: res.data.userId,
        authorName: `${res.data.firstName} ${res.data.lastName}`
      }))
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Không thể tải thông tin tác giả",
        variant: "destructive"
      })
    }
  }

  const handleImageUpload = async (e) => {
  const file = e.target.files[0];
  if (!file) return;

  const formData = new FormData();
  formData.append('file', file);

  setIsLoading(true);
  try {
    const response = await fetch('/api/upload', {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      throw new Error('Upload failed');
    }

    const data = await response.json();
    setFormData(prev => ({
      ...prev,
      image: data.url
    }));
    setPreviewImage(data.url);
    
    toast({
      title: "Thành công",
      description: "Tải ảnh lên thành công"
    });
  } catch (error) {
    toast({
      title: "Lỗi",
      description: "Không thể tải ảnh lên: " + error.message,
      variant: "destructive"
    });
  } finally {
    setIsLoading(false);
  }
};

  const handleSubmit = async (e) => {
    e.preventDefault()
    setIsLoading(true)

    try {
      if (!formData.image) {
        toast({
          title: "Lỗi",
          description: "Vui lòng tải lên một hình ảnh",
          variant: "destructive"
        })
        setIsLoading(false)
        return
      }

      const submissionData = {
        ...formData,
        image: formData.image // Use the uploaded image URL
      }

      if (id) {
        await mockNewsDataService.updateNews(id, submissionData)
        toast({
          title: "Thành công",
          description: "Bài viết đã được cập nhật",
        })
      } else {
        await mockNewsDataService.createNews(submissionData)
        toast({
          title: "Thành công",
          description: "Bài viết đã được tạo",
        })
      }
      router.push('/admin/news')
    } catch (error) {
      toast({
        title: "Lỗi",
        description: error.message,
        variant: "destructive"
      })
    } finally {
      setIsLoading(false)
    }
  }

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: value
    }))
  }

  return (
    <div className="container mx-auto pt-10 pl-64">
      {id ? 
        <h1 className="text-2xl font-semibold mb-4">Chỉnh sửa bài viết</h1> : 
        <h1 className="text-2xl font-semibold mb-4">Tạo Bài Viết Mới</h1>
      }

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <Label htmlFor="title">Tiêu Đề</Label>
            <Input
              id="title"
              name="title"
              value={formData.title}
              onChange={handleChange}
              placeholder="Nhập tiêu đề của bài đăng"
              required
              disabled={isLoading}
            />
          </div>

          <div>
            <Label htmlFor="description">Mô Tả</Label>
            <Textarea
              id="description"
              name="description"
              value={formData.description}
              onChange={handleChange}
              placeholder="Nhập mô tả hoặc tóm tắt bài đăng"
              required
              disabled={isLoading}
            />
          </div>

          <div>
            <Label htmlFor="content">Nội Dung</Label>
            <Textarea
              id="content"
              name="content"
              value={formData.content}
              onChange={handleChange}
              placeholder="Nhập chi tiết nội dung bài đăng"
              required
              className="h-40"
              disabled={isLoading}
            />
          </div>

          <div>
            <Label htmlFor="image">Ảnh bài viết</Label>
            <Input
              id="image"
              type="file"
              accept="image/*"
              onChange={handleImageUpload}
              disabled={isLoading}
              className="mt-2"
            />
          </div>

          {/* ...existing dropzone code... */}

          <Button 
            type="submit" 
            className="w-full bg-blue-500 hover:bg-blue-600"
            disabled={isLoading}
          >
            {isLoading ? 'Đang xử lý...' : (id ? 'Cập nhật bài' : 'Đăng Bài')}
          </Button>
        </form>

        <Card className="w-full">
          <CardHeader>
            <CardTitle>Xem Trước</CardTitle>
            <CardDescription>Tác giả: {formData.authorName}</CardDescription>
          </CardHeader>
          <CardContent>
            {previewImage ? (
              <div className="mt-2 relative h-48">
                <Image
                  src={previewImage}
                  alt="Preview"
                  fill
                  className="object-cover rounded-md"
                />
              </div>
            ) : (
              <div className="mt-2 h-48 flex items-center justify-center bg-gray-100 rounded-md">
                <p className="text-gray-500">Chưa có ảnh</p>
              </div>
            )}
            <h2 className="text-xl font-semibold mb-2">{formData.title || 'Tiêu Đề'}</h2>
            <p className="text-gray-600 mb-4">{formData.description || 'Mô tả'}</p>
            <div className="prose max-w-none">
              {formData.content ? (
                <p className="whitespace-pre-wrap">{formData.content}</p>
              ) : (
                <p className="text-gray-400">Nội dung sẽ xuất hiện tại đây</p>
              )}
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}