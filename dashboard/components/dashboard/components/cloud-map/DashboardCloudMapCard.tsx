import DashboardCloudMapChart from './DashboardCloudMapChart';
import DashboardCloudMapTooltip from './DashboardCloudMapTooltip';
import { DashboardCloudMapRegions } from './hooks/useCloudMap';
import useCloudMapTooltip from './hooks/useCloudMapTooltip';

type DashboardCloudMapCardProps = {
  data: DashboardCloudMapRegions | undefined;
};
function DashboardCloudMapCard({ data }: DashboardCloudMapCardProps) {
  const { tooltip, setTooltip } = useCloudMapTooltip();

  return (
    <div className="w-full rounded-lg bg-white py-4 px-6 pb-6">
      <div className="-mx-6 flex items-center justify-between border-b border-black-200/40 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-black-900">Cloud map</p>
          <div className="mt-1"></div>
          <p className="text-xs text-black-300">
            Analyze which regions have active resources
          </p>
        </div>
        <div className="h-[60px]"></div>
      </div>
      <div className="mt-8"></div>
      <DashboardCloudMapChart regions={data} setTooltip={setTooltip} />
      <DashboardCloudMapTooltip tooltip={tooltip} />
      <div className="flex gap-4 text-xs text-black-300">
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-info-600"></div>Active region
        </div>
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-black-300"></div>
          Inactive region
        </div>
      </div>
    </div>
  );
}

export default DashboardCloudMapCard;
