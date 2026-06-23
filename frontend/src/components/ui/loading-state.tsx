import { Skeleton } from "./skeleton"

type LoadingStateProps = {
  className?: string
  titleWidth?: string
  subtitleWidth?: string
  rows?: number
}

function LoadingState({
  className,
  titleWidth = "w-44",
  subtitleWidth = "w-72",
  rows = 3,
}: LoadingStateProps) {
  return (
    <div className={className}>
      <Skeleton className={`h-8 ${titleWidth}`} />
      <Skeleton className={`h-4 ${subtitleWidth}`} />
      <div className="space-y-3 pt-2">
        {Array.from({ length: rows }).map((_, index) => (
          <Skeleton key={index} className="h-10 w-full" />
        ))}
      </div>
    </div>
  )
}

export { LoadingState }
