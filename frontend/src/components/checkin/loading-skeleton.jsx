import React from "react";
import { Skeleton } from "@/components/ui/skeleton";

const LoadingSkeleton = () => {
  return (
    <div className="container mx-auto p-6 space-y-6">
      {/* Step Indicator Skeleton */}
      <div className="flex space-x-4">
        {Array.from({ length: 4 }).map((_, index) => (
          <Skeleton key={index} className="h-4 w-1/4 rounded" />
        ))}
      </div>

      {/* Content Skeleton */}
      <div className="space-y-4">
        <Skeleton className="h-6 w-48 rounded" />
        <Skeleton className="h-8 w-full rounded" />
        <Skeleton className="h-8 w-3/4 rounded" />
        <Skeleton className="h-8 w-1/2 rounded" />
      </div>

      {/* Table Skeleton */}
      <div className="space-y-2">
        {Array.from({ length: 5 }).map((_, index) => (
          <div key={index} className="flex space-x-2">
            <Skeleton className="h-6 w-1/6 rounded" />
            <Skeleton className="h-6 w-1/2 rounded" />
            <Skeleton className="h-6 w-1/4 rounded" />
          </div>
        ))}
      </div>

      {/* Button Skeleton */}
      <div className="flex space-x-4">
        <Skeleton className="h-10 w-32 rounded" />
        <Skeleton className="h-10 w-24 rounded" />
      </div>
    </div>
  );
};

export default LoadingSkeleton;
