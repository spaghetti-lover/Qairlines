import { Skeleton } from "@/components/ui/skeleton"
import { Button } from "@/components/ui/button"

function SkeletonFlightHeader() {
  return (
    <div className="p-4 bg-white shadow">
      <Skeleton className="h-8 w-48 mb-2" />
      <div className="flex gap-4">
        <Skeleton className="h-6 w-24" />
        <Skeleton className="h-6 w-24" />
        <Skeleton className="h-6 w-24" />
      </div>
    </div>
  )
}

function SkeletonFlightSideFilter() {
  return (
    <div className="w-64 p-4 bg-white shadow">
      <Skeleton className="h-6 w-32 mb-4" />
      {[1, 2, 3].map((i) => (
        <div key={i} className="mb-4">
          <Skeleton className="h-4 w-24 mb-2" />
          <Skeleton className="h-8 w-full" />
        </div>
      ))}
    </div>
  )
}

function SkeletonFlightCard() {
  return (
    <div className="bg-white p-4 shadow rounded-lg mb-4">
      <div className="flex justify-between mb-4">
        <Skeleton className="h-6 w-24" />
        <Skeleton className="h-6 w-24" />
      </div>
      <div className="flex justify-between">
        <Skeleton className="h-4 w-32" />
        <Skeleton className="h-4 w-32" />
      </div>
    </div>
  )
}

export default function Loading() {
  return (
    <div>
      <SkeletonFlightHeader />
      <div className="flex flex-col md:flex-row gap-4 p-4 bg-gray-100 min-h-screen max-w-6xl m-auto">
        <SkeletonFlightSideFilter />
        <div className="flex-1 space-y-4">
          {[1, 2, 3].map((i) => (
            <SkeletonFlightCard key={i} />
          ))}
          <Skeleton className="h-4 w-48 mx-auto" />
          <Button className="w-full bg-orange text-white" disabled>
            Quay lại trang chủ
          </Button>
        </div>
      </div>
    </div>
  )
}

