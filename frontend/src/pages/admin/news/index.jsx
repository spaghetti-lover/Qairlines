'use client'

import Link from 'next/link'
import { useEffect, useState } from 'react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from "@/components/ui/alert-dialog"
import { toast } from "@/hooks/use-toast"
import { Search, Plus, Edit, Trash2, Eye } from 'lucide-react'
import { useRouter } from 'next/router'
import mockNewsDataService from '@/pages/mockNewsData'

export default function PostManagementPage() {
  const router = useRouter()
  const [posts, setPosts] = useState([])
  const [searchTerm, setSearchTerm] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    getAllNews()
  }, [])

  const getAllNews = async () => {
    setIsLoading(true)
    try {
      const res = await mockNewsDataService.getAllNews()
      setPosts(res.data.map(item => ({
        id: item.id,
        title: item.title,
        author: item.authorName,
        description: item.description,
        createdAt: formatDate(item.createdAt)
      })))
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Không thể tải danh sách bài viết: " + error.message,
        variant: "destructive"
      })
    } finally {
      setIsLoading(false)
    }
  }

  const formatDate = (dateString) => {
    try {
      const date = new Date(dateString)
      return date.toLocaleString('vi-VN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit'
      }).replace(',', '')
    } catch (error) {
      return dateString
    }
  }

  const filteredPosts = posts.filter(post =>
    post.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
    post.author.toLowerCase().includes(searchTerm.toLowerCase())
  )

  const handleDeletePost = async (id) => {
    try {
      await mockNewsDataService.deleteNews(id)
      setPosts(posts.filter(post => post.id !== id))
      toast({
        title: "Thành công",
        description: "Bài viết đã được xóa thành công."
      })
    } catch (error) {
      toast({
        title: "Xóa bài viết không thành công",
        description: "Đã có lỗi xảy ra: " + error.message,
        variant: "destructive"
      })
    }
  }

  const handleRefresh = () => {
    getAllNews();
  }

  return (
    <div className="container mx-auto pt-10 pl-64 space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-semibold mb-4">Thông Tin & Bài Đăng</h1>
        <Button 
          onClick={handleRefresh} 
          variant="outline"
          disabled={isLoading}
        >
          {isLoading ? 'Đang tải...' : 'Làm mới'}
        </Button>
      </div>

      <div className="flex justify-between items-center">
        <div className="relative w-96">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Tìm kiếm theo tiêu đề hoặc tác giả..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="pl-8"
          />
        </div>

        <Link href="/admin/news/post" passHref>
          <Button className="bg-yellow-400 hover:bg-yellow-500 text-black font-medium">
            <Plus className="mr-2 h-4 w-4" />
            BÀI VIẾT MỚI
          </Button>
        </Link>
      </div>

      {/* Thống kê */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
        <div className="bg-blue-50 p-4 rounded-lg">
          <h3 className="text-sm font-medium text-gray-600">Tổng số bài viết</h3>
          <p className="text-2xl font-bold text-blue-600">{posts.length}</p>
        </div>
        <div className="bg-green-50 p-4 rounded-lg">
          <h3 className="text-sm font-medium text-gray-600">Kết quả tìm kiếm</h3>
          <p className="text-2xl font-bold text-green-600">{filteredPosts.length}</p>
        </div>
        <div className="bg-purple-50 p-4 rounded-lg">
          <h3 className="text-sm font-medium text-gray-600">Bài viết hôm nay</h3>
          <p className="text-2xl font-bold text-purple-600">
            {posts.filter(post => {
              const today = new Date().toDateString();
              const postDate = new Date(post.createdAt).toDateString();
              return today === postDate;
            }).length}
          </p>
        </div>
      </div>

      {isLoading ? (
        <div className="flex justify-center items-center py-8">
          <div className="text-gray-500">Đang tải danh sách bài viết...</div>
        </div>
      ) : (
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Tiêu đề</TableHead>
              <TableHead>Tác giả</TableHead>
              <TableHead>Ngày viết</TableHead>
              <TableHead>Tùy chỉnh</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredPosts.length === 0 ? (
              <TableRow>
                <TableCell colSpan={4} className="text-center py-8 text-gray-500">
                  {searchTerm ? 'Không tìm thấy bài viết nào phù hợp' : 'Chưa có bài viết nào'}
                </TableCell>
              </TableRow>
            ) : (
              filteredPosts.map((post) => (
                <TableRow key={post.id}>
                  <TableCell className="max-w-md">
                    <div className="truncate" title={post.title}>
                      {post.title}
                    </div>
                    {post.description && (
                      <div className="text-sm text-gray-500 truncate mt-1" title={post.description}>
                        {post.description}
                      </div>
                    )}
                  </TableCell>
                  <TableCell>{post.author}</TableCell>
                  <TableCell>{post.createdAt}</TableCell>
                  <TableCell>
                    <div className="flex space-x-2">
                      <Link href={`/admin/news/post?id=${post.id}`} passHref>
                        <Button variant="outline" size="icon" title="Chỉnh sửa">
                          <Edit className="h-4 w-4" />
                        </Button>
                      </Link>

                      <Link href={`/news/${post.id}`} passHref>
                        <Button variant="outline" size="icon" title="Xem bài viết">
                          <Eye className="h-4 w-4" />
                        </Button>
                      </Link>

                      <AlertDialog>
                        <AlertDialogTrigger asChild>
                          <Button variant="outline" size="icon" title="Xóa bài viết">
                            <Trash2 className="h-4 w-4" color="#ff0000"/>
                          </Button>
                        </AlertDialogTrigger>
                        <AlertDialogContent>
                          <AlertDialogHeader>
                            <AlertDialogTitle>Bạn có chắc chắn muốn xóa bài viết này không?</AlertDialogTitle>
                            <AlertDialogDescription>
                              Hành động này không thể hoàn tác. Bài viết "{post.title}" sẽ bị xóa vĩnh viễn khỏi hệ thống.
                            </AlertDialogDescription>
                          </AlertDialogHeader>
                          <AlertDialogFooter>
                            <AlertDialogCancel>Hủy bỏ</AlertDialogCancel>
                            <AlertDialogAction 
                              onClick={() => handleDeletePost(post.id)} 
                              className="bg-red-600 hover:bg-red-500"
                            >
                              Xóa bài viết
                            </AlertDialogAction>
                          </AlertDialogFooter>
                        </AlertDialogContent>
                      </AlertDialog>
                    </div>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      )}
    </div>
  )
}