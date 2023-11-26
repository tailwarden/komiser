function DashboardCloudMapSkeleton() {
  return (
    <div
      data-testid="loading"
      className="min-h-[396px] w-full animate-pulse rounded-lg bg-white px-6 py-4 pb-6"
    >
      <div className="-mx-6 flex items-center justify-between border-b border-gray-300 px-6 pb-4">
        <div>
          <div className="h-3 w-24 rounded-lg bg-cyan-200"></div>
          <div className="mt-2"></div>
          <div className="h-3 w-48 rounded-lg bg-cyan-200"></div>
        </div>
        <div className="h-[60px]"></div>
      </div>
      <div className="mt-8"></div>
      <div className="-mx-6 -ml-20 min-w-full">
        <picture>
          <img src="/assets/img/others/world.svg" alt="world map" />
        </picture>
      </div>
      <div className="mt-12"></div>
      <div className="flex gap-4 text-xs text-gray-500">
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-blue-500"></div>
          <div className="h-3 w-24 rounded-lg bg-cyan-200"></div>
        </div>
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-gray-500"></div>
          <div className="h-3 w-24 rounded-lg bg-cyan-200"></div>
        </div>
      </div>
    </div>
  );
}

export default DashboardCloudMapSkeleton;
