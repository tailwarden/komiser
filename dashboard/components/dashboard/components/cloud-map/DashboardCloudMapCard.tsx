import Button from '../../../button/Button';
import DashboardCloudMapChart from './DashboardCloudMapChart';
import DashboardCloudMapTooltip from './DashboardCloudMapTooltip';
import { DashboardCloudMapRegions } from './hooks/useCloudMap';
import useCloudMapExpand from './hooks/useCloudMapExpand';
import useCloudMapTooltip from './hooks/useCloudMapTooltip';

type DashboardCloudMapCardProps = {
  data: DashboardCloudMapRegions | undefined;
};
function DashboardCloudMapCard({ data }: DashboardCloudMapCardProps) {
  const { tooltip, setTooltip, sumOfResources } = useCloudMapTooltip({ data });
  const { isOpen, toggle } = useCloudMapExpand();

  return (
    <div
      data-testid="cloudMap"
      className={`${
        isOpen ? 'fixed inset-0 z-30 origin-left animate-scale' : ''
      } w-full rounded-lg bg-white px-6 py-4 pb-6`}
    >
      <div className="-mx-6 flex items-center justify-between border-b border-gray-300 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-gray-950">Cloud map</p>
          <div className="mt-1"></div>
          <p className="text-xs text-gray-500">
            Analyze which regions have active resources
          </p>
        </div>
        <div className="flex h-[60px] items-center">
          <Button style="ghost" size="sm" onClick={toggle}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="20"
              height="20"
              fill="none"
              viewBox="0 0 24 24"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="1.5"
                d="M21 9V3h-6M3 15v6h6M21 3l-7.5 7.5M10.5 13.5L3 21"
              ></path>
            </svg>
          </Button>
        </div>
      </div>
      <div className="mt-8"></div>
      {data && (
        <>
          <DashboardCloudMapChart
            regions={data}
            setTooltip={setTooltip}
            isOpen={isOpen}
          />
          <DashboardCloudMapTooltip
            tooltip={tooltip}
            sumOfResources={sumOfResources}
          />
        </>
      )}
      <div className="flex gap-4 text-xs text-gray-500">
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-blue-500"></div>Active region
        </div>
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-gray-500"></div>
          Inactive region
        </div>
      </div>
    </div>
  );
}

export default DashboardCloudMapCard;
