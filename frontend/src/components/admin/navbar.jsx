import Link from 'next/link'
import { useRouter } from 'next/router'
import { LayoutDashboard, Plane, CalendarDays, User, Users, UserPen, FileText, TicketCheck, MapPin } from 'lucide-react'

const navItems = [
  { 
    name: 'Tổng quan', 
    href: '/admin/dashboard', 
    icon: LayoutDashboard 
  },
  { 
    name: 'Quản lý chuyến bay', 
    href: '/admin/flights', 
    icon: CalendarDays 
  },
  {
    name: 'Quản lý tàu bay',
    href: '/admin/aircraft',
    icon: Plane
  },
  {
    name: 'Quản lý sân bay',
    href: '/admin/airports',
    icon: MapPin
  },
  { 
    name: 'Quản lý khách hàng', 
    href: '/admin/customers', 
    icon: Users 
  },
  { 
    name: 'Quản lý đặt vé', 
    href: '/admin/tickets', 
    icon: TicketCheck 
  },
  { 
    name: 'Thông tin & bài đăng', 
    href: '/admin/news', 
    icon: FileText 
  },
  { 
    name: 'Quản lý admin', 
    href: '/admin/members', 
    icon: UserPen 
  },
  { 
    name: 'Hồ sơ cá nhân', 
    href: '/admin/profile', 
    icon: User
  },
]

export default function Navbar() {
  const router = useRouter()
  return (
    <div className="h-screen w-64 bg-zinc-900 text-zinc-400">
      <div className="p-6 border-b border-zinc-800">
        <div className="flex items-center gap-3">
          <div className="p-1.5 bg-orange rounded">
            <Plane className="w-5 h-5 text-white" />
          </div>
          <h1 className="text-lg font-semibold text-white">QAIRLINE ADMIN</h1>
        </div>
      </div>
      <nav className="p-4">
        {navItems.map((item) => (
          <Link
            key={item.name}
            href={item.href}
            className={`flex items-center gap-3 px-4 py-3 rounded-lg text-sm transition-colors ${router.pathname === item.href 
              ? 'bg-orange/10 text-orange' 
              : 'text-white hover:bg-orange/10 hover:text-orange'
            }`}
          >
            <item.icon className="w-5 h-5" />
            <span>{item.name}</span>
          </Link>
        ))}
      </nav>
    </div>
  )
}