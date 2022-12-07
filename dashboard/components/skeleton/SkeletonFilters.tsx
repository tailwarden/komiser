function SkeletonFilters() {
  return (
    <div className="flex flex-wrap items-center gap-x-4 gap-y-2 bg-white py-2 px-6 rounded-lg mb-8 h-14">
      <div className="text-sm text-black-400">Filters</div>
      <div className="flex items-center gap-2 w-[20%] h-10 py-2 px-3 bg-komiser-100 rounded animate-pulse">
        <div className="bg-komiser-200/50 rounded-full h-5 w-5 flex-shrink-0"></div>
        <div className="bg-komiser-200/50 rounded h-5 w-full"></div>
      </div>
    </div>
  );
}

export default SkeletonFilters;
