import Button from '../../../button/Button';
import {
  CostExplorerQueryGranularityProps,
  CostExplorerQueryGroupProps
} from './hooks/useCostExplorer';

type DashboardCostExplorerErrorProps = {
  fetch: (
    provider?: CostExplorerQueryGroupProps,
    granularity?: CostExplorerQueryGranularityProps,
    startDate?: string,
    endDate?: string
  ) => void;
};

function DashboardCostExplorerError({
  fetch
}: DashboardCostExplorerErrorProps) {
  const rows: number[] = Array.from(Array(8).keys());

  return (
    <div className="w-full rounded-lg bg-white px-6 py-4 pb-6">
      <div className="-mx-6 flex items-center justify-center gap-6 border-b border-gray-300 px-6 pb-4">
        <p className="text-sm text-gray-700">
          There was an error loading the cost explorer.
        </p>
        <div className="flex-shrink-0">
          <Button style="secondary" size="sm" onClick={() => fetch()}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="none"
              viewBox="0 0 24 24"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="1.5"
                d="M22 12c0 5.52-4.48 10-10 10s-8.89-5.56-8.89-5.56m0 0h4.52m-4.52 0v5M2 12C2 6.48 6.44 2 12 2c6.67 0 10 5.56 10 5.56m0 0v-5m0 5h-4.44"
              ></path>
            </svg>
            Try again
          </Button>
        </div>
        <div className="h-[60px]"></div>
      </div>
      <div className="mt-8"></div>
      <div className="h-[22rem]">
        <table className="h-[90%] w-full table-auto">
          <tbody>
            {rows.map(idx => (
              <tr key={idx}>
                <td className="border border-background-ds"></td>
                <td className="border border-background-ds"></td>
                <td className="border border-background-ds"></td>
                <td className="border border-background-ds"></td>
              </tr>
            ))}
          </tbody>
        </table>
        <div className="mt-4"></div>
        <div className="flex justify-center gap-8">
          <div className="flex items-center gap-2">
            <div className="h-5 w-5 rounded-full bg-gray-300"></div>
            <div className="h-3 w-24 rounded-lg bg-gray-300"></div>
          </div>
          <div className="flex items-center gap-2">
            <div className="h-5 w-5 rounded-full bg-gray-300"></div>
            <div className="h-3 w-12 rounded-lg bg-gray-300"></div>
          </div>
          <div className="flex items-center gap-2">
            <div className="h-5 w-5 rounded-full bg-gray-300"></div>
            <div className="h-3 w-36 rounded-lg bg-gray-300"></div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default DashboardCostExplorerError;
