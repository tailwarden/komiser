function DashboardResourcesManagerSkeleton() {
  return (
    <div className="flex min-h-[360px] w-full animate-pulse flex-col gap-4 overflow-hidden rounded-lg bg-white py-4 px-6 pb-6">
      <div className=" -mx-6 flex items-center justify-between border-b border-black-200/40">
        <div className="px-6 pb-4">
          <div className="h-3 w-48 rounded-lg bg-komiser-200/50"></div>
          <div className="mt-2"></div>
          <div className="h-3 w-24 rounded-lg bg-komiser-200/50"></div>
        </div>
        <div className="px-6 pb-4">
          <div className="h-[60px] w-[9rem]"></div>
        </div>
      </div>
      <div className="h-[60px] w-full rounded-lg bg-komiser-200/50"></div>
      <div className="flex items-center justify-between px-6">
        <div className="min-h-[250px] min-w-[250px] rounded-full border-[50px] border-komiser-200/50 bg-white"></div>
        <div className="flex flex-col gap-4">
          <div className="flex items-center gap-2">
            <div className="h-4 w-4 rounded-lg bg-komiser-200/50"></div>
            <div className="h-3 w-24 rounded-lg bg-komiser-200/50"></div>
          </div>
          <div className="flex items-center gap-2">
            <div className="h-4 w-4 rounded-lg bg-komiser-200/50"></div>
            <div className="h-3 w-24 rounded-lg bg-komiser-200/50"></div>
          </div>
          <div className="flex items-center gap-2">
            <div className="h-4 w-4 rounded-lg bg-komiser-200/50"></div>
            <div className="h-3 w-24 rounded-lg bg-komiser-200/50"></div>
          </div>
          <div className="flex items-center gap-2">
            <div className="h-4 w-4 rounded-lg bg-komiser-200/50"></div>
            <div className="h-3 w-24 rounded-lg bg-komiser-200/50"></div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default DashboardResourcesManagerSkeleton;
