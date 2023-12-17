function SkeletonFilters() {
  return (
    <div className="mb-6 flex h-14 flex-wrap items-center gap-x-4 gap-y-2 rounded-lg bg-white px-6 py-2">
      <div className="text-sm text-gray-700">Filters</div>
      <div className="flex h-10 w-[20%] animate-pulse items-center gap-2 rounded bg-gray-50 px-3 py-2">
        <div className="h-5 w-5 flex-shrink-0 rounded-full bg-cyan-200"></div>
        <div className="h-5 w-full rounded bg-cyan-200"></div>
      </div>
    </div>
  );
}

export default SkeletonFilters;
