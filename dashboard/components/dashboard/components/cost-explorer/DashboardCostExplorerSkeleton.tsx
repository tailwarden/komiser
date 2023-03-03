function DashboardCostExplorerSkeleton() {
  const rows: number[] = Array.from(Array(8).keys());
  return (
    <div className="w-full animate-pulse rounded-lg bg-white py-4 px-6 pb-6">
      <div className="-mx-6 flex flex-wrap items-center justify-between border-b border-black-200/40">
        <div className="px-6 pb-4">
          <div className="h-3 w-24 rounded-lg bg-komiser-200/50"></div>
          <div className="mt-2"></div>
          <div className="h-3 w-36 rounded-lg bg-komiser-200/50"></div>
        </div>
        <div className="flex flex-col gap-4 px-6 pb-4 md:flex-row md:flex-wrap">
          <div className="h-[60px] w-[177.5px] rounded-lg bg-komiser-200/50"></div>
          <div className="h-[60px] w-[177.5px] rounded-lg bg-komiser-200/50"></div>
          <div className="h-[60px] w-[177.5px] rounded-lg bg-komiser-200/50"></div>
        </div>
      </div>
      <div className="mt-8"></div>
      <div className="h-[22rem]">
        <table className="h-[90%] w-full table-auto">
          <tbody>
            {rows.map(idx => (
              <tr key={idx}>
                <td className="border border-black-150"></td>
                <td className="border border-black-150"></td>
                <td className="border border-black-150"></td>
                <td className="border border-black-150"></td>
              </tr>
            ))}
          </tbody>
        </table>
        <div className="mt-4"></div>
        <div className="flex justify-center gap-8">
          <div className="flex items-center gap-2">
            <div className="h-5 w-5 rounded-full bg-komiser-200/50"></div>
            <div className="h-3 w-24 rounded-lg bg-komiser-200/50"></div>
          </div>
          <div className="flex items-center gap-2">
            <div className="h-5 w-5 rounded-full bg-komiser-200/50"></div>
            <div className="h-3 w-12 rounded-lg bg-komiser-200/50"></div>
          </div>
          <div className="flex items-center gap-2">
            <div className="h-5 w-5 rounded-full bg-komiser-200/50"></div>
            <div className="h-3 w-36 rounded-lg bg-komiser-200/50"></div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default DashboardCostExplorerSkeleton;
