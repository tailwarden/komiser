function SkeletonFilters() {
  return (
    <div className="mb-6 flex h-14 flex-wrap items-center gap-x-4 gap-y-2 rounded-lg bg-white py-2 px-6">
      <div className="text-sm text-black-400">Filters</div>
      <div className="flex h-10 w-[20%] animate-pulse items-center gap-2 rounded bg-komiser-100 py-2 px-3">
        <div className="h-5 w-5 flex-shrink-0 rounded-full bg-komiser-200/50"></div>
        <div className="h-5 w-full rounded bg-komiser-200/50"></div>
      </div>
    </div>
  );
}

export default SkeletonFilters;
