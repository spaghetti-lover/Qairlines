import { Skeleton } from "@/components/ui/skeleton"

export function SkeletonFlightCard() {
  return (
    <div className="bg-white p-4 shadow rounded-lg mb-4">
      <div className="flex justify-between mb-4">
        <Skeleton className="h-6 w-24" />
        <Skeleton className="h-6 w-24" />
      </div>
      <div className="flex justify-between mb-2">
        <Skeleton className="h-4 w-32" />
        <Skeleton className="h-4 w-32" />
      </div>
      <div className="flex justify-between">
        <Skeleton className="h-4 w-24" />
        <Skeleton className="h-4 w-24" />
      </div>
    </div>
  )
}

